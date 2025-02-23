package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	if authHeader == nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if len(authHeader) < 1 {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	bearerToken := authHeader[0]

	if bearerToken == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := strings.Split(bearerToken, " ")[1]

	id, isAdmin, err := m.jwt.ValidateToken(token)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
	}

	ctx.Locals("userId", id)
	ctx.Locals("isAdmin", isAdmin)
	return ctx.Next()
}
