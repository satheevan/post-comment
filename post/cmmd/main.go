package main

import (
	"post/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// Initialize default config
	app.Use(cors.New())

	//or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000/,http://localhost:3000",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	route.RoutePost(app)
	app.Listen(":3003")
}
