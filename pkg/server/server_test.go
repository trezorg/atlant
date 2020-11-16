package server

import (
	"context"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"testing"

	"github.com/trezorg/atlant/pkg/db"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/trezorg/atlant/internal/logger"
	"github.com/trezorg/atlant/pkg/loader"
	pb "github.com/trezorg/atlant/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const csvData = `test1;100
test2;102
test3;103`

type FakeLoader struct {
	*loader.CSVLoader
}

func (f FakeLoader) DownloadFile(url string) (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader(csvData)), nil
}

type FakeDatabase struct{}

func (s *FakeDatabase) Save(ctx context.Context, products []loader.Product) error {
	return nil
}
func (s *FakeDatabase) List(ctx context.Context, page *pb.Page) (*pb.Products, error) {
	if page.Limit == db.MaxLimit {
		return nil, status.Error(codes.InvalidArgument, "wrong limit specified")
	}
	return &pb.Products{}, nil
}

func (s *FakeDatabase) SaveChannel(
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
			if err := intercept(len(buf)); err != nil {
				return err
			}
			buf = buf[:0]
		}
	}
	if len(buf) > 0 {
		if err := intercept(len(buf)); err != nil {
			return err
		}
	}
	return nil
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	log := logger.InitDefaultLogger()

	fakeDB := &FakeDatabase{}
	fakeLoader := &FakeLoader{CSVLoader: loader.NewCsvLoader(log)}

	pb.RegisterAtlantServiceServer(server, NewWithLogger(fakeDB, fakeLoader, log))

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func Test_AtlantService_List_Handler(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)

	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewAtlantServiceClient(conn)

	requests := []*pb.Page{
		{
			Limit: 10,
			Sorting: &pb.Sorting{
				Field: pb.SortingField_PRICE,
				Order: pb.SortingOrder_DESC,
			},
			Cursor: &pb.Cursor{},
		},
		{
			Limit: db.MaxLimit + 100,
			Sorting: &pb.Sorting{
				Field: pb.SortingField_PRICE,
				Order: pb.SortingOrder_DESC,
			},
			Cursor: &pb.Cursor{
				Name:  "test",
				Field: "100",
			},
		},
	}

	t.Run("PageCorrectRequest", func(t *testing.T) {
		response, err := client.List(ctx, requests[0])
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Len(t, response.Products, 0)
		require.Nil(t, response.Cursor)
	})
	t.Run("PageInCorrectLimit", func(t *testing.T) {
		_, err := client.List(ctx, requests[1])
		pbErr, ok := status.FromError(err)
		require.True(t, ok)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, pbErr.Code())
	})

}

func Test_AtlantService_Fetch_Handler(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)

	defer func() {
		_ = conn.Close()
	}()

	client := pb.NewAtlantServiceClient(conn)

	t.Run("Fetch", func(t *testing.T) {
		stream, err := client.Fetch(ctx, &pb.FetchRequest{
			Url:        "http://example",
			SkipHeader: false,
			Separator:  ';',
		})
		require.NoError(t, err)
		require.NotNil(t, stream)
		var states []*pb.FetchState
		for {
			state, err := stream.Recv()
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
			states = append(states, state)
		}
		lines := strings.Split(csvData, "\n")
		require.Len(t, states, len(lines))
		require.NotNil(t, states)
		require.Equal(t, states[0].State, pb.StateEnum_IN_PROGRESS)
		require.Equal(t, int(states[0].LoadedRecords), 0)
		require.Equal(t, states[1].State, pb.StateEnum_IN_PROGRESS)
		require.Greater(t, int(states[1].LoadedRecords), 0)
		lastState := states[len(states)-1]
		require.Equal(t, lastState.State, pb.StateEnum_SUCCESS)
		require.Equal(t, int(lastState.LoadedRecords), len(lines))
	})

}
