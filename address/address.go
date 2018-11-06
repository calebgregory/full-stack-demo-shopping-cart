package address

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Type    string `json:"type";sql:"not null;type:ENUM('shipping', 'billing')"`
}

func New(db *gorm.DB) HttpHandler {
	return NewHttpHandler(NewServer(NewReaderWriter(db)))
}

type GetAllRequest struct{}

type GetAllResponse struct {
	Addresses []*Address `json:"addresses"`
	util.ErringResponse
}

type GetOneRequest struct {
	ID int `json:"id"`
}

type GetOneResponse struct {
	Address *Address `json:"address"`
	util.ErringResponse
}

type CreateRequest struct {
	Address *Address `json:"address"`
}

type CreateResponse struct {
	Address *Address `json:"address"`
	util.ErringResponse
}

type UpdateRequest struct {
	Address *Address `json:"address"`
}

type UpdateResponse struct {
	Address *Address `json:"address"`
	util.ErringResponse
}

type DeleteRequest struct {
	Address *Address `json:"address"`
}

type DeleteResponse struct {
	util.ErringResponse
}
