package customer

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/address"
	"github.com/calebgregory/full-stack-demo-shopping-cart/profile"
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	ProfileID int              `json:"profileID"`
	Profile   *profile.Profile `json:"profile"`
	AddressID int              `json:"addressID"`
	Address   *address.Address `json:"address"`
}

func New(db *gorm.DB) HttpHandler {
	return NewHttpHandler(NewServer(NewReaderWriter(db)))
}

type GetAllRequest struct{}

type GetAllResponse struct {
	Customers []*Customer `json:"customers"`
	util.ErringResponse
}

type GetOneRequest struct {
	ID int `json:"id"`
}

type GetOneResponse struct {
	Customer *Customer `json:"customer"`
	util.ErringResponse
}

type CreateRequest struct {
	Customer *Customer `json:"customer"`
}

type CreateResponse struct {
	Customer *Customer `json:"customer"`
	util.ErringResponse
}

type UpdateRequest struct {
	Customer *Customer `json:"customer"`
}

type UpdateResponse struct {
	Customer *Customer `json:"customer"`
	util.ErringResponse
}

type DeleteRequest struct {
	Customer *Customer `json:"customer"`
}

type DeleteResponse struct {
	util.ErringResponse
}
