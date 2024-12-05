package handlers

import "github.com/gofiber/fiber/v2"

func HandleError(e error, c *fiber.Ctx, status int) error {
	return c.Status(status).JSON(Error{
		e.Error(),
		status,
	})
}
