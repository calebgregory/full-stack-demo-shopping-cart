package product

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"log"
	"net/http"
)

type Server interface {
	GetAll(*GetAllRequest) *GetAllResponse
	Create(*CreateRequest) *CreateResponse
}

type Service struct {
	store ReaderWriter
}

func NewServer(store ReaderWriter) Server {
	return &Service{store: store}
}

func (s *Service) GetAll(req *GetAllRequest) (res *GetAllResponse) {
	res = &GetAllResponse{}
	ps, err := s.store.GetAll()
	if err != nil {
		log.Printf("service %s", err)
		res.Err = &util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		return
	}
	res.Products = ps
	return
}

func (s *Service) Create(req *CreateRequest) (res *CreateResponse) {
	res = &CreateResponse{}
	p, err := s.store.Create(req.Product)
	if err != nil {
		log.Printf("service %s", err)
		res.Err = &util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		return
	}
	res.Product = p
	return
}
