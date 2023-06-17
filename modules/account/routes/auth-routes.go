package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/lib/middleware"
	"github.com/thoriqadillah/linktrim/modules/account/controller"
)

func Routes(r fiber.Router) {
	router := r.Group("/auth")

	router.Get("/user", middleware.Auth, controller.GetUser)
	router.Post("/register", controller.Register)
	router.Post("/login", controller.Login)
}
