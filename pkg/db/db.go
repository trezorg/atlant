package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/trezorg/atlant/pkg/loader"
	pb "github.com/trezorg/atlant/pkg/proto"
)

const (
	DefaultLimit     = 100
	MaxLimit         = 1000
	MinLimit         = 1
	SortingOrderAsc  = 1
	SortingOrderDesc = -1
)

type SortingOrder int

func (s SortingOrder) Valid() bool {
	switch s {
	case SortingOrderDesc, SortingOrderAsc:
		return true
	}
	return false
}

type Database interface {
	SaveChannel(ctx context.Context, in chan loader.Record, bulkSize int, intercept func(int) error) error
	Save(ctx context.Context, products []loader.Product) error
	List(ctx context.Context, page *pb.Page) (*pb.Products, error)
}

func ToSortingFieldName(f pb.SortingField) string {
	switch f {
	case pb.SortingField_PRICE:
		return "price"
	case pb.SortingField_NAME:
		return "name"
	case pb.SortingField_PRICE_CHANGES:
		return "price_changes"
	case pb.SortingField_UPDATED_AT:
		return "updated_at"
	}
	return ""
}

func FromProductCursorField(field string, sortField pb.SortingField) (interface{}, error) {
	switch sortField {
	case pb.SortingField_NAME:
		return field, nil
	case pb.SortingField_PRICE, pb.SortingField_PRICE_CHANGES:
		return strconv.Atoi(field)
	case pb.SortingField_UPDATED_AT:
		return time.Parse(time.RFC3339, field)
	}
	return "", fmt.Errorf("unknown field: %s", field)
}

func ProductCursorField(product *pb.Product, field pb.SortingField) (string, error) {
	switch field {
	case pb.SortingField_NAME:
		return product.Name, nil
	case pb.SortingField_PRICE:
		return strconv.Itoa(int(product.Price)), nil
	case pb.SortingField_PRICE_CHANGES:
		return strconv.Itoa(int(product.PriceChanges)), nil
	case pb.SortingField_UPDATED_AT:
		tm, err := ptypes.Timestamp(product.UpdatedAt)
		if err != nil {
			return "", err
		}
		return tm.Format(time.RFC3339), nil
	}
	return "", fmt.Errorf("unknown field: %s", field)
}

func ToSortingOrder(s pb.SortingOrder) SortingOrder {
	switch s {
	case pb.SortingOrder_DESC:
		return SortingOrderDesc
	default:
		return SortingOrderAsc
	}
}
