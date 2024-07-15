package main

import (
	"context"
	"errors"
	"github.com/cemayan/searchengine/api/read"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/db"

	"github.com/cemayan/searchengine/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Init(constants.ReadApi)
	db.Init(constants.ReadApi)
}

func main() {
	server := read.NewServer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
