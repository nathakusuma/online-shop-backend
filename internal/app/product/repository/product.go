package repository

import (
	"gorm.io/gorm"
	"online-shop-backend/internal/domain/entity"
)

type ProductMySQLItf interface {
	GetAllProducts() ([]entity.Product, error)
	GetSpecificProduct(product *entity.Product) error
	CreateProduct(product entity.Product) error
	UpdateProduct(product entity.Product) error
	DeleteProduct(product entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (p *ProductMySQL) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := p.db.Find(&products).Error
	return products, err
}

func (p *ProductMySQL) GetSpecificProduct(product *entity.Product) error {
	return p.db.First(product).Error
}

func (p *ProductMySQL) CreateProduct(product entity.Product) error {
	return p.db.Create(&product).Error
}

func (p *ProductMySQL) UpdateProduct(product entity.Product) error {
	return p.db.Updates(&product).Error
}

func (p *ProductMySQL) DeleteProduct(product entity.Product) error {
	tx := p.db.Delete(&product)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return tx.Error
}
