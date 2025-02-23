package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authorization(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin")

	if isAdmin == false {
		return ctx.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return ctx.Next()
}
