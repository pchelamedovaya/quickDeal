package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ad struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Price       float64   `gorm:"not null"`
	AuthorID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Author User `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE"`
}

func (a *Ad) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
