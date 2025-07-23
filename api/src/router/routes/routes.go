package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// This represents all routes
type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// input all routes inside router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
