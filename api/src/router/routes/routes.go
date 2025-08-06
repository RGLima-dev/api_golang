package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}
	return r
}
