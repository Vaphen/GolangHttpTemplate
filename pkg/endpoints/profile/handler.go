package profile

import "net/http"
import "../../core"
import "../../utils"

type ProfileHandler struct {
	service core.Restable
}

func NewHandler(service core.Restable) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func (p ProfileHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	utils.Success(w, p.service.Get())
}

func (p ProfileHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(203)
	utils.Success(w, p.service.Post(nil))
}

func (p ProfileHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(202)
	utils.Success(w, p.service.Put(nil))
}

func (p ProfileHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	utils.Success(w, p.service.Delete(nil))
}
