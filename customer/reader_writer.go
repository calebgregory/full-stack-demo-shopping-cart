package customer

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ReaderWriter interface {
	GetAll() ([]*Customer, error)
	GetOne(id int) (*Customer, error)
	Create(*Customer) (*Customer, error)
	Update(*Customer) (*Customer, error)
	Delete(*Customer) error
}

type Store struct {
	db *gorm.DB
}

func NewReaderWriter(db *gorm.DB) ReaderWriter {
	db.AutoMigrate(&Customer{})
	return &Store{db: db}
}

func (s *Store) GetAll() ([]*Customer, error) {
	var ps []*Customer
	if err := s.db.Find(&ps).Error; err != nil {
		return ps, errors.Wrap(err, "store")
	}
	return ps, nil
}

func (s *Store) GetOne(id int) (*Customer, error) {
	var p Customer
	if err := s.db.First(&p, id).Error; err != nil {
		return &p, errors.Wrap(err, "store")
	}
	return &p, nil
}

func (s *Store) Create(p *Customer) (*Customer, error) {
	if err := s.db.Create(p).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}
	return p, nil
}

func (s *Store) Update(update *Customer) (r *Customer, err error) {
	var p Customer

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

func (s *Store) Delete(p *Customer) (err error) {
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
