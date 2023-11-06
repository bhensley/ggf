package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) PerformRedirect(c *fiber.Ctx) error {
	redirect, err := h.redirectStore.GetByShortURL(c.Params("shortURL"))

	// Make sure nothing bad happened
	if err != nil || redirect == nil {
		return c.Render("BadShortURL", fiber.Map{
			"Title": "URL Not Found",
		})
	}

	// If it's not nil, it has a URL, so send them off!
	return c.Redirect(redirect.OriginalURL, fiber.StatusMovedPermanently)
}

func (h *Handler) DisplayAnalytics(c *fiber.Ctx) error {
	return c.Render("Analytics", fiber.Map{
		"Title": "Analytics Coming Soon",
	})
}
