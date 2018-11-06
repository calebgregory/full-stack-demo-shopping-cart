package product

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ReaderWriter interface {
	GetAll() ([]*Product, error)
	GetOne(id int) (*Product, error)
	Create(*Product) (*Product, error)
	Update(*Product) (*Product, error)
	Delete(*Product) error
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

func (s *Store) Update(update *Product) (r *Product, err error) {
	var p Product

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.First(&p).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}

	if err = tx.Model(&p).Update(update).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}

	if err = tx.Commit().Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}

	return &p, nil
}

func (s *Store) Delete(p *Product) (err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.First(p).Error; err != nil {
		return errors.Wrap(err, "store")
	}

	if err = tx.Delete(p).Error; err != nil {
		return errors.Wrap(err, "store")
	}

	if err = tx.Commit().Error; err != nil {
		return errors.Wrap(err, "store")
	}

	return nil
}
