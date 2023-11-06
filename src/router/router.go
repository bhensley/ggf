package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func New() *fiber.App {
	engine := html.New("../client/views", ".html")
	fiber := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	engine.Reload(true)

	fiber.Static("/assets", "../client/assets")
	fiber.Use(logger.New())
	fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, POST",
	}))

	return fiber
}
