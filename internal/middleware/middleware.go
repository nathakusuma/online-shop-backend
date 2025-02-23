package middleware

import (
	"github.com/gofiber/fiber/v2"
	"online-shop-backend/internal/infra/jwt"
)

type MiddlewareItf interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt *jwt.JWT
}

func NewMiddleware(jwt *jwt.JWT) MiddlewareItf {
	return &Middleware{
		jwt: jwt,
	}
}
