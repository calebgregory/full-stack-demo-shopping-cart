package profile

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Name string `json:"name"`
}

func New(db *gorm.DB) HttpHandler {
	return NewHttpHandler(NewServer(NewReaderWriter(db)))
}

type GetAllRequest struct{}

type GetAllResponse struct {
	Profiles []*Profile `json:"profiles"`
	util.ErringResponse
}

type GetOneRequest struct {
	ID int `json:"id"`
}

type GetOneResponse struct {
	Profile *Profile `json:"profile"`
	util.ErringResponse
}

type CreateRequest struct {
	Profile *Profile `json:"profile"`
}

type CreateResponse struct {
	Profile *Profile `json:"profile"`
	util.ErringResponse
}

type UpdateRequest struct {
	Profile *Profile `json:"profile"`
}

type UpdateResponse struct {
	Profile *Profile `json:"profile"`
	util.ErringResponse
}

type DeleteRequest struct {
	Profile *Profile `json:"profile"`
}

type DeleteResponse struct {
	util.ErringResponse
}
