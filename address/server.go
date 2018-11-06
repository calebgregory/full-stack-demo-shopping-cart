package address

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/pkg/errors"
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
	res.Addresses = ps
	return
}

func (s *Service) GetOne(req *GetOneRequest) (res *GetOneResponse) {
	res = &GetOneResponse{}

	if req.ID == 0 {
		res.Err = util.NewResponseError(
			errors.New("bad data"),
			"No ID found on request body; format { id: Int }",
			400,
		)
		return
	}

	p, err := s.store.GetOne(req.ID)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Address = p
	return
}

func (s *Service) Create(req *CreateRequest) (res *CreateResponse) {
	res = &CreateResponse{}

	if req.Address == nil {
		res.Err = util.NewResponseError(
			errors.New("bad data"),
			"Address not found on request body; format { address: {...} }",
			400,
		)
		return
	}

	p, err := s.store.Create(req.Address)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Address = p
	return
}

func (s *Service) Update(req *UpdateRequest) (res *UpdateResponse) {
	res = &UpdateResponse{}

	if req.Address == nil || req.Address.ID == 0 {
		res.Err = util.NewResponseError(
			errors.New("bad data"),
			"no address with ID found on request body; format { address: { id, ... } }",
			400,
		)
		return
	}

	p, err := s.store.Update(req.Address)
	if err != nil {
		res.Err = util.NewResponseError(err)
		return
	}
	res.Address = p
	return
}

func (s *Service) Delete(req *DeleteRequest) (res *DeleteResponse) {
	res = &DeleteResponse{}

	if req.Address == nil || req.Address.ID == 0 {
		res.Err = util.NewResponseError(
			errors.New("bad data"),
			"no address with ID found on request body; format { address: { id, ... } }",
			400,
		)
		return
	}

	err := s.store.Delete(req.Address)
	if err != nil {
		res.Err = util.NewResponseError(err)
	}
	return
}
