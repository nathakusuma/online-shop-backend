package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"online-shop-backend/internal/app/product/repository"
	"online-shop-backend/internal/domain/dto"
	"online-shop-backend/internal/domain/entity"
)

type ProductUsecaseItf interface {
	GetProducts() ([]dto.ResponseProduct, error)
	GetSpecificProduct(id uuid.UUID) (dto.ResponseProduct, error)
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseProduct, error)
	UpdateProduct(id uuid.UUID, request dto.RequestUpdateProduct) error
	DeleteProduct(id uuid.UUID) error
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (p *ProductUsecase) GetProducts() ([]dto.ResponseProduct, error) {
	products, err := p.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	resp := make([]dto.ResponseProduct, len(products))
	for i, product := range products {
		resp[i] = product.ParseToDTO()
	}

	return resp, nil
}

func (p *ProductUsecase) GetSpecificProduct(id uuid.UUID) (dto.ResponseProduct, error) {
	product := entity.Product{
		ID: id,
	}
	if err := p.ProductRepository.GetSpecificProduct(&product); err != nil {
		return dto.ResponseProduct{}, err
	}

	return product.ParseToDTO(), nil
}

func (p *ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseProduct, error) {
	product := entity.Product{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		PhotoURL:    request.PhotoURL,
	}

	if err := p.ProductRepository.CreateProduct(product); err != nil {
		return dto.ResponseProduct{}, err
	}

	return product.ParseToDTO(), nil
}

func (p *ProductUsecase) UpdateProduct(id uuid.UUID, request dto.RequestUpdateProduct) error {
	product := entity.Product{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		PhotoURL:    request.PhotoURL,
	}

	if err := p.ProductRepository.UpdateProduct(product); err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecase) DeleteProduct(id uuid.UUID) error {
	product := entity.Product{
		ID: id,
	}

	if err := p.ProductRepository.DeleteProduct(product); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Product not found")
		}
		return err
	}

	return nil
}
