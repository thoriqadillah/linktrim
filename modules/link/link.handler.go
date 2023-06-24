package link

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/lib/helper"
	"github.com/thoriqadillah/linktrim/modules/account/model"
)

var storer = NewStore(db.DB())

func createLink(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*model.User)
	var payload linkCreate

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(
				fmt.Sprintf("Error parsing request body: %s", err.Error()),
			))
	}

	payload.Owner = user.ID
	if err := storer.Create(c.Context(), payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(
				fmt.Sprintf("Error creating link: %s", err.Error()),
			))
	}

	return c.Status(http.StatusCreated).
		JSON(helper.SuccessResponse(1))
}

func getLinks(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*model.User)

	links, err := storer.GetAll(c.Context(), user.ID, helper.Paginate(c))
	if err != nil {
		return c.Status(http.StatusNotFound).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(links))
}

func getOneLink(c *fiber.Ctx) error {
	id := c.Params("id")
	linkID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(
				fmt.Sprintf("Error parsing UUID: %s", err.Error()),
			))
	}

	link, err := storer.GetOne(c.Context(), linkID)
	if err != nil {
		return c.Status(http.StatusNotFound).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(link))
}
func updateLink(c *fiber.Ctx) error {
	id := c.Params("id")
	linkID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(
				fmt.Sprintf("Error parsing UUID: %s", err.Error()),
			))
	}

	var payload linkUpdate
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	if err := storer.Update(c.Context(), linkID, payload); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(id))
}

func deleteLink(c *fiber.Ctx) error {
	id := c.Params("id")
	linkID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(helper.ErrorResponse(
				fmt.Sprintf("Error parsing UUID: %s", err.Error()),
			))
	}

	if err := storer.Delete(c.Context(), linkID); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(helper.ErrorResponse(err.Error()))
	}
	return c.Status(http.StatusOK).
		JSON(helper.SuccessResponse(id))
}

func trimmedRedirect(c *fiber.Ctx) error {
	trimmed := c.Params("trimmed")
	link, err := storer.GetOriginalFromTrimmed(c.Context(), trimmed)
	if err != nil {
		return c.Status(http.StatusNotFound).
			JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Redirect(link.Original, http.StatusTemporaryRedirect)
}
