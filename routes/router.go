package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/modules/auth"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	auth.Routes(v1)

}
