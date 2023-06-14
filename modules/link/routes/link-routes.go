package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/lib/middleware"
)

func Routes(r fiber.Router) {
	router := r.Group("/link", middleware.Auth)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hemlo")
	})
}
