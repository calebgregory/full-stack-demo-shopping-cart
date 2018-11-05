package order

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"log"
	"net/http"
)

type HttpHandler interface {
	HandleAddProduct(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	server Server
}

func NewHttpHandler(server Server) HttpHandler {
	return &Handler{server: server}
}

func (h *Handler) HandleAddProduct(w http.ResponseWriter, r *http.Request) {
	var req AddProductRequest
	if err := util.BindJSON(r, &req); err != nil {
		util.WriteFailedResponse(w, r, err)
		return
	}

	res := h.server.AddProduct(&req)

	util.WriteResponse(w, res)
}
