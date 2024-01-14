package route

import (
	"github.com/gofiber/fiber/v2"
	controller "orchestrtorservices.com/orchestrator/controller/AuthenticationController"
)

func RouteAuthUser(app *fiber.App) {
	ResgisterUser := controller.AuthenticationUser{}
	app.Post("/register", ResgisterUser.Register)
}
