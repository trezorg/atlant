package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/trezorg/atlant/internal/logger"
	"github.com/trezorg/atlant/pkg/db/mongo"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"github.com/trezorg/atlant/pkg/db"
	"github.com/trezorg/atlant/pkg/loader"
	pb "github.com/trezorg/atlant/pkg/proto"
)

type AtlantServer struct {
	pb.UnimplementedAtlantServiceServer
	log    *logrus.Entry
	db     db.Database
	loader loader.Loader
}

func (a *AtlantServer) Fetch(req *pb.FetchRequest, stream pb.AtlantService_FetchServer) error {
	if err := stream.Send(&pb.FetchState{
		State:         pb.StateEnum_IN_PROGRESS,
		LoadedRecords: 0,
	}); err != nil {
		return err
	}
	a.log.Infof("Processing req %s...", req.Url)
	reader, err := a.loader.DownloadFile(req.Url)
	if err != nil {
		return err
	}
	records, err := a.loader.ReadCSV(reader, req.Separator, req.SkipHeader)
	if err != nil {
		return err
	}
	ctx, done := context.WithCancel(context.Background())
	defer done()
	processedCount := 0
	bulkSize := 100
	if err := a.db.SaveChannel(ctx, records, bulkSize, func(n int) error {
		processedCount += n
		if err := stream.Send(&pb.FetchState{
			State:         pb.StateEnum_IN_PROGRESS,
			LoadedRecords: int32(processedCount),
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return stream.Send(&pb.FetchState{
		State:         pb.StateEnum_SUCCESS,
		LoadedRecords: int32(processedCount),
	})
}

func (a *AtlantServer) List(ctx context.Context, page *pb.Page) (*pb.Products, error) {
	SetPageDefaults(page)
	return a.db.List(ctx, page)
}

func NewWithLogger(db db.Database, loader loader.Loader, logger *logrus.Entry) *AtlantServer {
	return &AtlantServer{log: logger, db: db, loader: loader}
}

func Start(ctx context.Context, connectionURL, host string, port int, logLevel string) error {
	log := logger.InitLogger(logLevel, false)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	log.Infof("Initializing mongodb %s...", connectionURL)
	mongoDB, err := mongo.New(ctx, connectionURL)
	if err != nil {
		return err
	}

	listenerAddress := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", listenerAddress)
	if err != nil {
		return err
	}

	var grpcServer *grpc.Server
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		var opts []grpc.ServerOption
		grpcServer = grpc.NewServer(opts...)
		pb.RegisterAtlantServiceServer(
			grpcServer,
			NewWithLogger(mongoDB, loader.NewCsvLoader(log), log),
		)
		log.Infof("Starting grpc server on %s...", listenerAddress)
		if err := grpcServer.Serve(listener); err != nil {
			log.Errorf("cannot start grpc server: %v", err)
		}
	}(&wg)

	select {
	case sig := <-stop:
		log.Infof("Caught sig: %+v. Waiting process is being stopped...", sig)
	case <-ctx.Done():
	}

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	wg.Wait()

	return nil
}
