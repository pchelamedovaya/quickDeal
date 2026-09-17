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

func (r *AdGormRepository) FindAll() ([]entity.Ad, error) {
	var ads []entity.Ad
	err := r.db.Preload("Author").Order("created_at desc").Find(&ads).Error
	return ads, err
}
