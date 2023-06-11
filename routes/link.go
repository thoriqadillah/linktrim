package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/modules/link"
)

func linkRoutes(r fiber.Router) {
	r.Post("/", link.CreateLink)
}
