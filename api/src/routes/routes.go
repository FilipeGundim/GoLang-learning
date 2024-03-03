package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Func                  func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
