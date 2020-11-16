package testing

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ory/dockertest/v3/docker"

	"github.com/trezorg/atlant/pkg/db/mongo"

	pb "github.com/trezorg/atlant/pkg/proto"

	"google.golang.org/grpc"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"github.com/trezorg/atlant/pkg/server"
)

const (
	csvData = `test1;100
test2;400
test3;100
test3;200`
	port = "27017"
	host = "127.0.0.1"
)

var mongoURI = fmt.Sprintf("mongodb://%s:%s", host, port)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "mongo",
		Tag:          "latest",
		ExposedPorts: []string{port},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(port): {{HostIP: host, HostPort: port}},
		},
	},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	ctx := context.Background()
	if err := pool.Retry(func() error {
		var err error
		_, err = mongo.New(ctx, mongoURI)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to mongodb: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func Test_AtlantService(t *testing.T) {

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(csvData))
	}))
	defer backend.Close()

	ctx, done := context.WithCancel(context.Background())
	host, port, logLevel, mongoURI := "0.0.0.0", 10000, "DEBUG", mongoURI

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		err := server.Start(ctx, mongoURI, host, port, logLevel)
		require.NoError(t, err)
	}()

	time.Sleep(time.Duration(1) * time.Second)

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", "127.0.0.1", port), grpc.WithInsecure())
	require.NoError(t, err)
	client := pb.NewAtlantServiceClient(conn)

	stream, err := client.Fetch(ctx, &pb.FetchRequest{
		Url:        backend.URL,
		SkipHeader: false,
		Separator:  ';',
	})
	require.NoError(t, err)

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
	require.Len(t, states, 3+len(lines)/100)

	db, err := mongo.New(ctx, mongoURI)
	require.NoError(t, err)

	records, err := db.List(ctx, server.SetPageDefaults(&pb.Page{
		Limit:   10,
		Sorting: nil,
		Cursor:  nil,
	}))
	require.NoError(t, err)
	require.Len(t, records.Products, 3)
	require.Equal(t, "test1", records.Products[0].Name)

	records, err = db.List(ctx, server.SetPageDefaults(&pb.Page{
		Limit: 10,
		Sorting: &pb.Sorting{
			Field: pb.SortingField_NAME,
			Order: pb.SortingOrder_DESC,
		},
		Cursor: nil,
	}))
	require.NoError(t, err)
	require.Len(t, records.Products, 3)
	require.Equal(t, "test3", records.Products[0].Name)

	records, err = db.List(ctx, server.SetPageDefaults(&pb.Page{
		Limit: 1,
		Sorting: &pb.Sorting{
			Field: pb.SortingField_PRICE,
			Order: pb.SortingOrder_DESC,
		},
		Cursor: nil,
	}))
	require.NoError(t, err)
	require.NotNil(t, records)
	require.Len(t, records.Products, 1)
	require.Equal(t, "test2", records.Products[0].Name)
	require.Equal(t, 400, int(records.Products[0].Price))
	require.NotNil(t, records.Cursor)

	var recs = records.Products

	for len(records.Products) > 0 {
		records, err = db.List(ctx, server.SetPageDefaults(&pb.Page{
			Limit: 1,
			Sorting: &pb.Sorting{
				Field: pb.SortingField_PRICE,
				Order: pb.SortingOrder_DESC,
			},
			Cursor: records.Cursor,
		}))
		require.NoError(t, err)
		require.NotNil(t, records)
		recs = append(recs, records.Products...)
	}

	require.Len(t, recs, 3, recs)

	done()

	wg.Wait()

}
