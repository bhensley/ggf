package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) DisplayHome(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Home",
	}, "_layout")
}
