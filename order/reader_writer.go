package order

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/jinzhu/gorm"
	"github.com/satchelhealth/errors"
	"log"
)

type ReaderWriter interface {
	AddProduct(*Order, *product.Product) (*Order, error)
}

type Store struct {
	db *gorm.DB
}

func NewReaderWriter(db *gorm.DB) ReaderWriter {
	db.AutoMigrate(&Order{})
	return &Store{db: db}
}

func (s *Store) AddProduct(order *Order, product *product.Product) (*Order, error) {
	if err := s.db.FirstOrCreate(order, order).Error; err != nil {
		return nil, errors.Wrap(err, "store first or create")
	}

	if err := s.db.Find(product, product).Error; err != nil {
		return nil, errors.Wrap(err, "store find product")
	}

	if err := s.db.Model(order).Association("Products").Append(product).Error; err != nil {
		return nil, errors.Wrap(err, "store association append")
	}

	return order, nil
}
