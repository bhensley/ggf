package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	r.Get("/", h.DisplayHome)
	r.Get("/:shortURL", h.PerformRedirect)
	r.Post("/:shortURL", h.CreateRedirect)
	r.Get("/:shortURL!", h.DisplayAnalytics)
}
