package repository

import (
	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetProducts() string
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (p *ProductMySQL) GetProducts() string {
	return "Get All Products Success"
}
