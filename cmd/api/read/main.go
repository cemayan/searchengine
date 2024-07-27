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
	pb "github.com/cemayan/searchengine/protos/event"
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

func subscribeToNats() {

	ws := service.WriteService{ProjectName: constants.ReadApi}

	subscribe := messaging.MessagingServer.Subscribe(constants.NatsEventsStream, "consumer-event")

	subscribe.Consume(func(msg jetstream.Msg) {
		var evt pb.Event
		proto.Unmarshal(msg.Data(), &evt)

		ws.Start(string(evt.Data))
		msg.Ack()
	})
}

func main() {

	server := read.NewServer()

	subscribeToNats()

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
