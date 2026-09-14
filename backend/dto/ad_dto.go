package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAdRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gte=0"`
}

type AdResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
}
