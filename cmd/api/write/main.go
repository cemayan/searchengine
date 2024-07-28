package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/cemayan/searchengine/api/write"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/messaging"
	"github.com/cemayan/searchengine/internal/service"
	pb "github.com/cemayan/searchengine/protos/backendreq"
	eventpb "github.com/cemayan/searchengine/protos/event"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configPath string // get config yaml path as flag
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path of config yaml")
	flag.Parse()
	// config initializer
	config.Init(constants.WriteApi, configPath)
	// db initializer
	db.Init(constants.WriteApi)
	// messaging initializer
	messaging.Init(constants.WriteApi)
}

type server struct {
	pb.UnimplementedDbServiceServer
}

func (s server) SendRequest(ctx context.Context, request *pb.BackendRequest) (*pb.BackendRequest, error) {

	svc := service.NewWriteService(constants.WriteApi)
	seErr := svc.AddRecordMetadataToDb(request)
	if seErr != nil {

		go func() {
			svc.PublishErrorsToNats(constants.NatsErrorsStream, seErr)

		}()

		return nil, errors.New("an error occurred while adding record metadata to db")
	}

	marshal, err := proto.Marshal(request)
	if err != nil {
		logrus.Errorln("an error occurred while marshalling request:", err)
		return nil, err
	}

	go func() {
		svc.PublishToNats(marshal, constants.NatsEventsStream, eventpb.EventType_RECORDMETADATA_CREATED, eventpb.EntityType_RecordMetadata)
	}()

	return &pb.BackendRequest{Items: request.Items}, nil
}

func backendGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetConfig(constants.WriteApi).Scraper.Server.Port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDbServiceServer(s, &server{})
	logrus.Printf("grpc server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	go func() {
		backendGrpcServer()
	}()

	writeServer := write.NewServer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		writeServer.Configure()
		if err := writeServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error starting server: %s\n", err)
		}
	}()

	logrus.Infof("Server started on port: %v\n", config.GetConfig(constants.WriteApi).Serve.Port)

	<-done
	logrus.Infoln("Server stopped")

	defer cancel()

	if err := writeServer.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown Failed:%+v", err)
	}

}
