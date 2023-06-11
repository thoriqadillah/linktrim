package auth

import "github.com/gofiber/fiber/v2"

func Routes(r fiber.Router) {
	router := r.Group("/auth")

	router.Post("/register", register)
	router.Post("/login", login)
}
