package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/lib/cache"
	"github.com/thoriqadillah/linktrim/lib/helper"
	"github.com/thoriqadillah/linktrim/lib/security"
	"github.com/thoriqadillah/linktrim/modules/account/model"
	"github.com/thoriqadillah/linktrim/modules/account/store"
)

var storer = store.NewStore(db.DB())

func Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return c.Status(http.StatusUnauthorized).
			JSON(helper.ErrorResponse("Unauthorized"))
	}

	token := strings.Replace(authHeader, "Bearer ", "", -1)
	userID, err := security.DecodeJWT(token)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusUnauthorized).
			JSON(helper.ErrorResponse("Unauthorized"))
	}

	var user *model.User
	cache := cache.NewUserCache()
	res, err := cache.Get(c.Context(), userID.String())
	if err != nil {
		if user, _ = storer.GetUser(c.Context(), userID); user == nil {
			return c.Status(http.StatusUnauthorized).
				JSON(helper.ErrorResponse("Unauthorized"))
		}
	}

	user = res.(*model.User)
	ctx := context.WithValue(c.Context(), "user", user)
	c.SetUserContext(ctx)

	return c.Next()
}
