package repository

import (
	"quickdeal/entity"

	"gorm.io/gorm"
)

type AdGormRepository struct {
	db *gorm.DB
}

func NewAdRepository(db *gorm.DB) *AdGormRepository {
	return &AdGormRepository{db: db}
}

func (r *AdGormRepository) Create(ad *entity.Ad) error {
	return r.db.Create(ad).Error
}
