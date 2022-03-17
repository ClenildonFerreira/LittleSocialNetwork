package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI                   string
	Method                string
	Riquisitions          func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

// Configure to put all routes inside the ROUTER
func Config(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, router := range routes {
		r.HandleFunc(router.URI, router.Riquisitions).Methods(router.Method)
	}

	return r
}
