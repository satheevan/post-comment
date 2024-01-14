package postcontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/controller"
	"orchestrtorservices.com/orchestrator/database"
	"orchestrtorservices.com/orchestrator/model"
	"orchestrtorservices.com/orchestrator/response"
)

type PostGetController struct {
	controller.BaseController
}

func (pg PostGetController) GetPostController(c *fiber.Ctx) error {
	fmt.Println("Orchestrator - Inside get all post")
	//Define
	var data []model.Post

	Db := database.ConnectPostDB()

	query := Db.Raw("SELECT * FROM Post")

	if query.Scan(&data); query.Error != nil {
		panic(query.Error)
	}
	return c.JSON(response.Response{Status: 200, Message: "Received all data", Data: map[string]interface{}{"data": &data}})
}
