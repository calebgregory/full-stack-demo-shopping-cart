package profile

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ReaderWriter interface {
	GetAll() ([]*Profile, error)
	GetOne(id int) (*Profile, error)
	Create(*Profile) (*Profile, error)
	Update(*Profile) (*Profile, error)
	Delete(*Profile) error
}

type Store struct {
	db *gorm.DB
}

func NewReaderWriter(db *gorm.DB) ReaderWriter {
	db.AutoMigrate(&Profile{})
	return &Store{db: db}
}

func (s *Store) GetAll() ([]*Profile, error) {
	var ps []*Profile
	if err := s.db.Find(&ps).Error; err != nil {
		return ps, errors.Wrap(err, "store")
	}
	return ps, nil
}

func (s *Store) GetOne(id int) (*Profile, error) {
	var p Profile
	if err := s.db.First(&p, id).Error; err != nil {
		return &p, errors.Wrap(err, "store")
	}
	return &p, nil
}

func (s *Store) Create(p *Profile) (*Profile, error) {
	if err := s.db.Create(p).Error; err != nil {
		return nil, errors.Wrap(err, "store")
	}
	return p, nil
}

func (s *Store) Update(update *Profile) (r *Profile, err error) {
	var p Profile

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

func (s *Store) Delete(p *Profile) (err error) {
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
