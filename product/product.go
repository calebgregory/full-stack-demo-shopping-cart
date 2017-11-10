package product

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Description string
	Price       float32
}

func New(db *gorm.DB) HttpHandler {
	return NewHttpHandler(NewServer(NewReaderWriter(db)))
}

type GetAllRequest struct{}

type GetAllResponse struct {
	Products []*Product `json:"products"`
	util.ErringResponse
}

type CreateRequest struct {
	Product *Product `json:"product"`
}

type CreateResponse struct {
	Product *Product `json:"product"`
	util.ErringResponse
}
