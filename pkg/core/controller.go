package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	handler RestHandler
	router *mux.Router
}

func NewController(handler RestHandler, router *mux.Router) *Controller {
	return &Controller{
		handler: handler,
		router: router,
	}
}

func (c Controller) Start(path string, pathparam string) {
	resourcePath := fmt.Sprintf("%s/{%s}", path, pathparam)
	c.router.Path(path).Methods(http.MethodPost).HandlerFunc(c.handler.PostHandler)
	c.router.Path(resourcePath).Methods(http.MethodGet).HandlerFunc(c.handler.GetHandler)
	c.router.Path(resourcePath).Methods(http.MethodDelete).HandlerFunc(c.handler.DeleteHandler)
	c.router.Path(resourcePath).Methods(http.MethodPut).HandlerFunc(c.handler.PutHandler)
}

