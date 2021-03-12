package client

import (
	"../core"
	"../endpoints/profile"
	"github.com/gorilla/mux"
	"net/http"
)

var routes = map[string]core.RestHandler {
	"/profiles": profile.NewHandler(profile.NewService()),
}

func Run(address string) {
	http.ListenAndServe(address, initRoutes(mux.Router{}))
}

func initRoutes(router mux.Router) *mux.Router {
	for path, handler := range routes {
		core.NewController(handler, &router).Start(path, "id")
	}
	return &router
}