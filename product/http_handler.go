package product

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"log"
	"net/http"
)

type HttpHandler interface {
	HandleGetAll(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	server Server
}

func NewHttpHandler(s Server) HttpHandler {
	return &Handler{s}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	var req GetAllRequest
	if err := util.BindJSON(r, &req); err != nil {
		log.Printf("product handler handle get all bind json %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := h.server.GetAll(&req)

	util.WriteResponse(w, res)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := util.BindJSON(r, &req); err != nil {
		log.Printf("product handler handle create bind json %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res := h.server.Create(&req)

	util.WriteResponse(w, res)
}
