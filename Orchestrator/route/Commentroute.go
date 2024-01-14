package route

import (
	"github.com/gofiber/fiber/v2"
	controller "orchestrtorservices.com/orchestrator/controller/CommentContoller"
)

func RouteComment(app *fiber.App) {
	CommentController := &controller.CommentController{}

	app.Post("/createcomment", CommentController.CreateCommentController)
	// app.Get("/deletecomment")
}
