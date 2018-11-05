package product

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
)

type Server interface {
	GetAll(*GetAllRequest) *GetAllResponse
	GetOne(*GetOneRequest) *GetOneResponse
	Create(*CreateRequest) *CreateResponse
	Update(*UpdateRequest) *UpdateResponse
	Delete(*DeleteRequest) *DeleteResponse
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
		res.Err = util.NewResponseError(err)
		return
	}
	res.Products = ps
	return
}

func (s *Service) GetOne(req *GetOneRequest) (res *GetOneResponse) {
	res = &GetOneResponse{}
	p, err := s.store.GetOne(req.ID)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Product = p
	return
}

func (s *Service) Create(req *CreateRequest) (res *CreateResponse) {
	res = &CreateResponse{}
	p, err := s.store.Create(req.Product)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Product = p
	return
}

func (s *Service) Update(req *UpdateRequest) (res *UpdateResponse) {
	res = &UpdateResponse{}
	p, err := s.store.Update(req.Product)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Product = p
	return
}

func (s *Service) Delete(req *DeleteRequest) (res *DeleteResponse) {
	res = &DeleteResponse{}
	err := s.store.Delete(req.Product)
	if err != nil {
		res.Err = util.NewResponseError(err)
	}
	return
}
