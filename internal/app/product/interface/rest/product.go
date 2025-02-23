package rest

import (
	"github.com/gofiber/fiber/v2"
	"online-shop-backend/internal/app/product/usecase"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup fiber.Router, productUsecase usecase.ProductUsecaseItf) {
	ProductHandler := ProductHandler{
		ProductUsecase: productUsecase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("", ProductHandler.GetAllProducts)
}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	return ctx.SendString(h.ProductUsecase.GetProducts())
}
