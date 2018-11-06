package address

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ReaderWriter interface {
	GetAll() ([]*Address, error)
	GetOne(id int) (*Address, error)
	Create(*Address) (*Address, error)
	Update(*Address) (*Address, error)
	Delete(*Address) error
}

type Store struct {
	db *gorm.DB
}

func NewReaderWriter(db *gorm.DB) ReaderWriter {
	db.AutoMigrate(&Address{})
	return &Store{db: db}
}

func (s *Store) GetAll() ([]*Address, error) {
	var ps []*Address
	if err := s.db.Find(&ps).Error; err != nil {
		return ps, errors.Wrap(err, "store")
	}
	return ps, nil
}

func (s *Store) GetOne(id int) (*Address, error) {
	var p Address
	if err := s.db.First(&p, id).Error; err != nil {
		return &p, errors.Wrap(err, "store")
	}
	return &p, nil
}

func (s *Store) Create(p *Address) (*Address, error) {
	if err := s.db.Create(p).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}
	return p, nil
}

func (s *Store) Update(update *Address) (r *Address, err error) {
	var p Address

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

func (s *Store) Delete(p *Address) (err error) {
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
