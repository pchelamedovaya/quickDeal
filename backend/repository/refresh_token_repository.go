package repository

import (
	"quickdeal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenGormRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenGormRepository {
	return &RefreshTokenGormRepository{db: db}
}

func (r *RefreshTokenGormRepository) Create(token *entity.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenGormRepository) FindByTokenHash(tokenHash string) (*entity.RefreshToken, error) {
	var token entity.RefreshToken
	err := r.db.Where("token_hash = ?", tokenHash).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenGormRepository) Revoke(id uuid.UUID) error {
	return r.db.Model(&entity.RefreshToken{}).Where("id = ?", id).Update("revoked", true).Error
}
