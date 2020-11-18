package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/golang/protobuf/ptypes"

	"github.com/sirupsen/logrus"
	pb "github.com/trezorg/atlant/pkg/proto"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func list(c *cli.Context) error {
	serverURI := c.String("server-uri")
	sorting := c.String("sorting")
	sortingOrder := c.String("order")
	var pbSorting pb.SortingField
	switch sorting {
	case "name":
		pbSorting = pb.SortingField_NAME
	case "price":
		pbSorting = pb.SortingField_PRICE
	case "date":
		pbSorting = pb.SortingField_UPDATED_AT
	default:
		return fmt.Errorf("unknown sorting field: Allowed: name, price, date")
	}
	var pbOrder pb.SortingOrder
	switch sortingOrder {
	case "asc":
		pbOrder = pb.SortingOrder_ASC
	case "desc":
		pbOrder = pb.SortingOrder_DESC
	default:
		return fmt.Errorf("unknown sorting order: Allowed: asc, desc")
	}
	ctx, done := context.WithCancel(context.Background())
	defer done()

	conn, err := grpc.DialContext(ctx, serverURI, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewAtlantServiceClient(conn)

	result, err := client.List(ctx, &pb.Page{Limit: 10, Sorting: &pb.Sorting{
		Field: pbSorting,
		Order: pbOrder,
	}})
	if err != nil {
		return err
	}

	for _, record := range result.Products {
		updatedAt, _ := ptypes.Timestamp(record.UpdatedAt)
		fmt.Printf(
			"%s => %d. Price changed: %d. Updated at: %s\n",
			record.Name,
			record.Price,
			record.PriceChanges,
			updatedAt,
		)
	}

	return nil

}

func fetch(c *cli.Context) error {
	serverURI := c.String("server-uri")
	fetchURI := c.String("fetch-uri")
	skipHeader := c.Bool("skip-header")
	separator := c.String("separator")
	ctx, done := context.WithCancel(context.Background())
	defer done()

	conn, err := grpc.DialContext(ctx, serverURI, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewAtlantServiceClient(conn)

	sep := []rune(separator)

	stream, err := client.Fetch(ctx, &pb.FetchRequest{
		Url:        fetchURI,
		SkipHeader: skipHeader,
		Separator:  sep[0],
	})
	if err != nil {
		return err
	}
	for {
		state, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s => %d\n", state.State, state.LoadedRecords)
	}
	return nil

}

func main() {
	app := prepareArgs()
	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("Cannot initialize application: %+v", err)
		os.Exit(1)
	}
}
