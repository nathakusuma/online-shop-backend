package entity

import (
	"github.com/google/uuid"
	"online-shop-backend/internal/domain/dto"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int8      `gorm:"type:smallint;not null"`
	PhotoURL    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p Product) ParseToDTO() dto.ResponseProduct {
	return dto.ResponseProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		PhotoURL:    p.PhotoURL,
	}
}
