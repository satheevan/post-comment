package route

import (
	"github.com/gofiber/fiber/v2"
	controller "orchestrtorservices.com/orchestrator/controller/PostContoller"
)

func RoutePost(app *fiber.App) {
	//PostController struct
	postController := &controller.PostController{}
	postDeleteController := &controller.PostDeleteContoller{}

	app.Post("/createpost", postController.CreatePostController)
	app.Post("/deletepost", postDeleteController.DeletePostController)
}
