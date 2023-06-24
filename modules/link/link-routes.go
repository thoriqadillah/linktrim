package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/lib/middleware"
)

func Routes(r fiber.Router) {
	r.Get("/:trimmed", trimmedRedirect)

	router := r.Group("/link", middleware.Auth)

	router.Get("/", getLinks)
	router.Get("/:id", getOneLink)
	router.Post("/", createLink)
	router.Put("/:id", updateLink)
	router.Delete("/:id", deleteLink)
}
