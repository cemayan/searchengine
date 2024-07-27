package write

import (
	"context"
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//go:generate oapi-codegen --config cfg.yaml api.yaml

// Server struct
type Server struct {
	server *http.Server
	router *Router
}

func (srv *Server) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func corsHandler() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: config.GetConfig(constants.WriteApi).Web.AllowedOrigins,
		AllowedMethods: []string{"GET", "HEAD", "POST", "DELETE"},
	})
}

func (srv *Server) Configure() {

	// get an `http.Handler` that we can use
	h := HandlerFromMux(srv, srv.router.router)
	srv.router.negroni.UseHandler(h)

	srv.server.Handler = corsHandler().Handler(srv.router.negroni)
}

func (srv *Server) ListenAndServe() error {
	// And we serve HTTP until the world ends.
	return srv.server.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown Failed:%+v", err)
	}

	return nil
}

func NewServer() Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.GetConfig(constants.WriteApi).Serve.Port),
		WriteTimeout: time.Second * 300,
		ReadTimeout:  time.Second * 300,
		IdleTimeout:  time.Second * 300,
	}
	return Server{server: server, router: NewRouter()}
}
