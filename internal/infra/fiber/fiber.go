package fiber

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

const idleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	return app
}
