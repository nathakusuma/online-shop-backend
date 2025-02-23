package dto

import "github.com/google/uuid"

type RequestCreateProduct struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3,max=100"`
	Price       int64  `json:"price" validate:"required,min=1000"`
	Stock       int8   `json:"stock" validate:"required,min=1"`
	PhotoURL    string `json:"photo_url" validate:"required,url"`
}

type ResponseProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Stock       int8      `json:"stock"`
	PhotoURL    string    `json:"photo_url"`
}

type RequestUpdateProduct struct {
	Title       string `json:"title" validate:"omitempty,min=3,max=50"`
	Description string `json:"description" validate:"omitempty,min=3,max=100"`
	Price       int64  `json:"price" validate:"omitempty,min=1000"`
	Stock       int8   `json:"stock" validate:"omitempty,min=1"`
	PhotoURL    string `json:"photo_url" validate:"omitempty,url"`
}
