package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/lib/middleware"
)

func Routes(r fiber.Router) {
	router := r.Group("/link", middleware.Auth)

	router.Get("/", GetLinks)
	router.Get("/:id", GetOneLink)
	router.Post("/", CreateLink)
	router.Put("/:id", UpdateLink)
	router.Delete("/:id", DeleteLink)
}
