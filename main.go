package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/thoriqadillah/linktrim/db"
	_ "github.com/thoriqadillah/linktrim/lib/env"
	"github.com/thoriqadillah/linktrim/routes"
)

func main() {
	db.Open()
	defer db.Close()

	app := fiber.New()
	app.Use(logger.New())

	routes.Setup(app)

	app.Listen(":8000")
}
