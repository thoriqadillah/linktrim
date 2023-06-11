package helper

import "github.com/gofiber/fiber/v2"

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		"message": "success",
		"data":    data,
	}
}

func ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		"message": message,
	}
}
