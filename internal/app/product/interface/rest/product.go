package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"online-shop-backend/internal/app/product/usecase"
	"online-shop-backend/internal/domain/dto"
	"online-shop-backend/internal/middleware"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecaseItf
	Middleware     middleware.MiddlewareItf
	Validator      *validator.Validate
}

func NewProductHandler(
	routerGroup fiber.Router,
	productUsecase usecase.ProductUsecaseItf,
	middleware middleware.MiddlewareItf,
	validator *validator.Validate,
) {
	ProductHandler := ProductHandler{
		ProductUsecase: productUsecase,
		Middleware:     middleware,
		Validator:      validator,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Use(ProductHandler.Middleware.Authentication)

	routerGroup.Get("", ProductHandler.GetAllProducts)
	routerGroup.Post("",
		ProductHandler.Middleware.Authorization,
		ProductHandler.CreateProduct)
	routerGroup.Get("/:id", ProductHandler.GetSpecificProduct)
	routerGroup.Patch("/:id",
		ProductHandler.Middleware.Authorization,
		ProductHandler.UpdateProduct)
	routerGroup.Delete("/:id",
		ProductHandler.Middleware.Authorization,
		ProductHandler.DeleteProduct)
}

func (h *ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, "Invalid UUID")
	}

	res, err := h.ProductUsecase.GetSpecificProduct(productID)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Get Specific Product Success",
		"payload": res,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res, err := h.ProductUsecase.GetProducts()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Get All Products Success",
		"payload": res,
	})
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestCreateProduct
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := h.Validator.Struct(request); err != nil {
		return err
	}

	res, err := h.ProductUsecase.CreateProduct(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Create Product Success",
		"payload": res,
	})
}

func (h *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestUpdateProduct
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := h.Validator.Struct(request); err != nil {
		return err
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, "Invalid UUID")
	}

	if err := h.ProductUsecase.UpdateProduct(productID, request); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Update Product Success",
	})
}

func (h *ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, "Invalid UUID")
	}

	if err := h.ProductUsecase.DeleteProduct(productID); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Delete Product Success",
	})
}
