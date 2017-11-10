package order

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Products []*product.Product `gorm:"many2many:order_products;"`
}

func New(db *gorm.DB) HttpHandler {
	return NewHttpHandler(NewServer(NewReaderWriter(db)))
}

type AddProductRequest struct {
	Order   *Order           `json:"order"`
	Product *product.Product `json:"product"`
}

type AddProductResponse struct {
	Order *Order `json:"order"`
	util.ErringResponse
}
