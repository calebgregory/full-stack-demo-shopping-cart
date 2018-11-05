package product

import (
	"github.com/jinzhu/gorm"
	"github.com/satchelhealth/errors"
)

type ReaderWriter interface {
	GetAll() ([]*Product, error)
	GetOne(id int) (*Product, error)
	Create(*Product) (*Product, error)
}

type Store struct {
	db *gorm.DB
}

func NewReaderWriter(db *gorm.DB) ReaderWriter {
	db.AutoMigrate(&Product{})
	return &Store{db: db}
}

func (s *Store) GetAll() ([]*Product, error) {
	var ps []*Product
	if err := s.db.Find(&ps).Error; err != nil {
		return ps, errors.Wrap(err, "store")
	}
	return ps, nil
}

func (s *Store) GetOne(id int) (*Product, error) {
	var p Product
	if err := s.db.First(&p, id).Error; err != nil {
		return &p, errors.Wrap(err, "store")
	}
	return &p, nil
}

func (s *Store) Create(p *Product) (*Product, error) {
	if err := s.db.Create(p).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}
	return p, nil
}
