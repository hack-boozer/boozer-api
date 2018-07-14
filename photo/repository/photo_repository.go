package repository

import (
	model "github.com/hack-boozer/boozer-api/photo"
	"github.com/jinzhu/gorm"
)

type photoRepository struct {
	Conn *gorm.DB
}

// NewPhotoRepository mount photo repository
func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		Conn: db,
	}
}

// PhotoRepository photo repository interface
type PhotoRepository interface {
	Create(photo *model.Photo) (*model.Photo, error)
}

func (m *photoRepository) Create(photo *model.Photo) (*model.Photo, error) {
	err := m.Conn.Create(photo).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}
