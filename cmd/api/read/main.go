package main

import (
	"context"
	"errors"
	"flag"
	"github.com/cemayan/searchengine/api/read"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/messaging"
	"github.com/cemayan/searchengine/internal/service"
	backendreqpb "github.com/cemayan/searchengine/protos/backendreq"
	eventpb "github.com/cemayan/searchengine/protos/event"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
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
	config.Init(constants.ReadApi, configPath)
	// db initializer
	db.Init(constants.ReadApi)
	//messaging initializer
	messaging.Init(constants.ReadApi)

}

func subscribeToNatsEvents() {

	ws := service.WriteService{ProjectName: constants.ReadApi}

	subscribe := messaging.MessagingServer.Subscribe(constants.NatsEventsStream, "consumer-event")

	subscribe.Consume(func(msg jetstream.Msg) {
		var event eventpb.Event
		proto.Unmarshal(msg.Data(), &event)

		if event.EntityType == eventpb.EntityType_Record {

			for _, err := range ws.Write(string(event.Data)) {

				go func() {
					ws.PublishErrorsToNats(constants.NatsErrorsStream, &err)
				}()

			}

		} else if event.EntityType == eventpb.EntityType_RecordMetadata {

			var br backendreqpb.BackendRequest
			proto.Unmarshal(event.Data, &br)

			if err := ws.AddRecordMetadataToDb(&br); err != nil {

				logrus.Errorln("Failed to add metadata to db:", err)

				go func() {
					ws.PublishErrorsToNats(constants.NatsErrorsStream, err)
				}()
			}
		}

		msg.Ack()
	})
}

func main() {

	server := read.NewServer()

	subscribeToNatsEvents()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		server.Configure()
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error starting server: %s\n", err)
		}
	}()

	logrus.Infof("Server started on port: %v\n", config.GetConfig(constants.ReadApi).Serve.Port)

	<-done
	logrus.Infoln("Server stopped")

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown Failed:%+v", err)
	}

}
