package order

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"log"
	"net/http"
)

type Server interface {
	AddProduct(*AddProductRequest) *AddProductResponse
}

type Service struct {
	store ReaderWriter
}

func NewServer(store ReaderWriter) Server {
	return &Service{store: store}
}

func (s *Service) AddProduct(req *AddProductRequest) (res *AddProductResponse) {
	res = &AddProductResponse{}
	order, err := s.store.AddProduct(req.Order, req.Product)
	if err != nil {
		log.Printf("order service %s", err)
		res.Err = &util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		return res
	}
	res.Order = order
	return
}
