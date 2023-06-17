package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/modules/account/model"
	"github.com/thoriqadillah/linktrim/modules/account/store"

	"github.com/thoriqadillah/linktrim/lib/cache"
	"github.com/thoriqadillah/linktrim/lib/helper"
	"github.com/thoriqadillah/linktrim/lib/security"
)

var storer = store.NewStore(db.DB())

func GetUser(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*model.User)
	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(user))
}

func Register(c *fiber.Ctx) error {
	var user model.UserCreate
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	if err := storer.Create(c.Context(), user); err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusCreated).
		JSON(helper.SuccessResponse(1))
}

func Login(c *fiber.Ctx) error {
	var payload model.UserLogin
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	user, err := storer.Login(c.Context(), payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	token := security.EncodeJWT(user.ID.String())

	toCache, err := json.Marshal(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	cache := cache.NewUserCache()
	if err := cache.Set(c.Context(), user.ID.String(), toCache); err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(token))
}
