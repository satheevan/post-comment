package route

import (
	"post/controller"

	"github.com/gofiber/fiber/v2"
)

func RoutePost(app *fiber.App) {
	app.Post("/api/createpost", controller.CreatePostController)
	app.Get("/api/getpost", controller.GetallPostController)
}

// Home
// app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, world ! this is post")
// 	})
