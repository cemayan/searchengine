package write

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Router struct {
	router  *mux.Router
	negroni *negroni.Negroni
}

func NewRouter() *Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/v1").Subrouter()
	n := negroni.Classic() // Includes some default middlewares

	return &Router{router: api, negroni: n}
}
