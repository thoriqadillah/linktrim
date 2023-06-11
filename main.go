package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db.Setup()
	routes.Setup(app)

	app.Listen(":8000")

	defer db.DB().Close()
}
