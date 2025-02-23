package usecase

import "online-shop-backend/internal/app/product/repository"

type ProductUsecaseItf interface {
	GetProducts() string
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (p *ProductUsecase) GetProducts() string {
	return p.ProductRepository.GetProducts()
}
