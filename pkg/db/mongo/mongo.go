package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/trezorg/atlant/pkg/db"
	"github.com/trezorg/atlant/pkg/loader"
	pb "github.com/trezorg/atlant/pkg/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	dbName         = "atlant"
	collectionName = "products"
	nameFieldName  = "name"
)

var (
	indexes = []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{
					Key:   nameFieldName,
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().SetUnique(true),
		},
		// pagination sort by price
		{
			Keys: bsonx.Doc{
				{
					Key:   "price",
					Value: bsonx.Int32(1),
				},
				{
					Key:   nameFieldName,
					Value: bsonx.Int32(1),
				},
			},
		},
		// pagination sort by updated_at
		{
			Keys: bsonx.Doc{
				{
					Key:   "updated_at",
					Value: bsonx.Int32(1),
				},
				{
					Key:   nameFieldName,
					Value: bsonx.Int32(1),
				},
			},
		},
		// pagination sort by price_changes
		{
			Keys: bsonx.Doc{
				{
					Key:   "price_changes",
					Value: bsonx.Int32(1),
				},
				{
					Key:   nameFieldName,
					Value: bsonx.Int32(1),
				},
			},
		},
	}
)

func isMongoDuplicateError(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code != 11000 {
				return false
			}
		}
	}
	return true
}

type DB struct {
	*mongo.Client
}

func (d *DB) ensureIndexes(ctx context.Context) error {
	_, err := d.products().Indexes().CreateMany(ctx, indexes)
	return err
}

func (d *DB) products() *mongo.Collection {
	return d.Client.Database(dbName).Collection(collectionName)
}

func sortOptions(page *pb.Page) bson.M {
	order := db.ToSortingOrder(page.Sorting.Order)
	field := db.ToSortingFieldName(page.Sorting.Field)
	if page.Sorting.Field == pb.SortingField_NAME {
		return bson.M{field: order}
	}
	return bson.M{
		field:         order,
		nameFieldName: order,
	}
}

func filterOptions(page *pb.Page) bson.M {
	if page.Cursor.Field == "" && page.Cursor.Name == "" {
		return bson.M{}
	}
	cursorField, err := db.FromProductCursorField(page.Cursor.Field, page.Sorting.Field)
	if err != nil {
		return bson.M{}
	}
	comparator := "$gt"
	if page.Sorting.Order == pb.SortingOrder_DESC {
		comparator = "$lt"
	}
	sortingField := db.ToSortingFieldName(page.Sorting.Field)
	if page.Sorting.Field == pb.SortingField_NAME {
		return bson.M{sortingField: bson.M{comparator: cursorField}}
	}
	return bson.M{
		"$or": bson.A{
			bson.M{
				sortingField: bson.M{comparator: cursorField},
			},
			bson.M{
				sortingField:  page.Cursor.Field,
				nameFieldName: bson.M{comparator: page.Cursor.Name},
			},
		},
	}
}

func New(ctx context.Context, connectionURL string) (*DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	_db := &DB{client}
	return _db, _db.ensureIndexes(ctx)
}

type mongoProduct struct {
	Name         string    `bson:"name"`
	Price        int       `bson:"price"`
	PriceChanges int       `bson:"price_changes"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func decodeProduct(c *mongo.Cursor) (*pb.Product, error) {
	p := mongoProduct{}
	if err := c.Decode(&p); err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.Product{
		Name:         p.Name,
		Price:        int32(p.Price),
		UpdatedAt:    ts,
		PriceChanges: int32(p.PriceChanges),
	}, nil
}

func (d *DB) List(ctx context.Context, page *pb.Page) (*pb.Products, error) {
	opts := options.Find()
	opts.SetLimit(int64(page.Limit))
	opts.SetSort(sortOptions(page))
	c, err := d.products().Find(ctx, filterOptions(page), opts)
	if err != nil {
		return nil, err
	}
	var products []*pb.Product
	for c.Next(ctx) {
		product, err := decodeProduct(c)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	cursor := pb.Cursor{}
	if len(products) > 0 {
		lastProduct := products[len(products)-1]
		if cursor.Field, err = db.ProductCursorField(lastProduct, page.Sorting.Field); err != nil {
			return nil, err
		}
		cursor.Name = lastProduct.Name
	}
	return &pb.Products{
		Products: products,
		Cursor:   &cursor,
	}, nil
}

func (d *DB) Save(ctx context.Context, products []loader.Product) error {
	if len(products) == 0 {
		return nil
	}
	operations := make([]mongo.WriteModel, 0, len(products))
	for _, product := range products {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{
			nameFieldName: product.Name,
			"price": bson.M{
				"$ne": product.Price,
			},
		})
		operation.SetUpdate(bson.M{
			"$set": bson.M{
				"price":      product.Price,
				"updated_at": time.Now().UTC(),
			},
			"$inc": bson.M{"price_changes": 1},
		})
		operation.SetUpsert(true)
		operations = append(operations, operation)
	}
	ordered := false
	_, err := d.products().BulkWrite(ctx, operations, &options.BulkWriteOptions{
		Ordered: &ordered,
	})
	// skip duplicated error
	if err != nil && !isMongoDuplicateError(err) {
		return err
	}
	return nil
}

func (d *DB) SaveChannel(
	ctx context.Context,
	in chan loader.Record,
	bulkSize int,
	intercept func(int) error,
) error {
	buf := make([]loader.Product, 0, bulkSize)
	for rec := range in {
		if rec.Err != nil {
			return rec.Err
		}
		buf = append(buf, rec.Product)
		if len(buf) == bulkSize {
			if err := d.Save(ctx, buf); err != nil {
				return err
			}
			if err := intercept(len(buf)); err != nil {
				return err
			}
			buf = buf[:0]
		}
	}
	if len(buf) > 0 {
		if err := d.Save(ctx, buf); err != nil {
			return err
		}
		if err := intercept(len(buf)); err != nil {
			return err
		}
	}
	return nil
}
