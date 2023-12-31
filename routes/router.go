package routes

import (
	"github.com/gofiber/fiber/v2"

	account "github.com/thoriqadillah/linktrim/modules/account/routes"
	"github.com/thoriqadillah/linktrim/modules/link"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	account.Routes(v1)
	link.Routes(v1)
}
