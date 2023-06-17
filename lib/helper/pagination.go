package helper

import "github.com/gofiber/fiber/v2"

type Pagination struct {
	Limit int
	Page  int
}

func Paginate(c *fiber.Ctx) Pagination {
	limit := c.QueryInt("limit", 25)
	page := c.QueryInt("page", 1)
	offset := limit * (page - 1)

	return Pagination{
		Limit: limit,
		Page:  offset,
	}
}
