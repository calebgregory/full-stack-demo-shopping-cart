package order

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CustomerID  uint
	SubmittedAt time.Time
	Products    []*OrderProduct `gorm:"ForeignKey:OrderID"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
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
