package main

import (
	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/database"
	"orchestrtorservices.com/orchestrator/route"
)

//local port value/

// func helloMessage(c *fiber.Ctx) {
// 	c.SendString("Orchestrator service")
// }

// func RouteSetUp(app *fiber.App) {
// 	app.Get("/", helloMessage)
// }

// Main function
func main() {

	app := fiber.New()
	// ========Database===========
	database.ConnectPostDB()
	database.ConnectCommentDB()

	// ========Route===========
	// RouteSetUp(app)

	//authentication Route
	route.RouteAuthUser(app)

	//Post
	route.RoutePost(app)

	//Comment
	route.RouteComment(app)

	app.Listen(":3000")
}
