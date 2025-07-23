package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// generate a router with all routers configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
