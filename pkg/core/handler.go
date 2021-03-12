package core

import "net/http"

type RestHandler interface {
	GetHandler
	PostHandler
	PutHandler
	DeleteHandler
}

type GetHandler interface{
	GetHandler(w http.ResponseWriter, r *http.Request)
}

type PostHandler interface{
	PostHandler(w http.ResponseWriter, r *http.Request)
}

type PutHandler interface{
	PutHandler(w http.ResponseWriter, r *http.Request)
}

type DeleteHandler interface{
	DeleteHandler(w http.ResponseWriter, r *http.Request)
}