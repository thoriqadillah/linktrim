package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/lib/helper"
)

var store = NewStore(db.DB())

func register(c *fiber.Ctx) error {
	var user userCreate
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	if err := store.Create(c.Context(), user); err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusCreated).
		JSON(helper.SuccessResponse(1))
}

func login(c *fiber.Ctx) error {
	var payload userLogin
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	user, err := store.Login(c.Context(), payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(user))
}
