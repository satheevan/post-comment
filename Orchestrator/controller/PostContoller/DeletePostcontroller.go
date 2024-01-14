package postcontroller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/controller"
	"orchestrtorservices.com/orchestrator/database"
	"orchestrtorservices.com/orchestrator/model"
	"orchestrtorservices.com/orchestrator/response"
)

type PostDeleteContoller struct {
	controller.BaseController
}

func (pc PostDeleteContoller) DeletePostController(c *fiber.Ctx) error {
	data := &model.Post{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedAt = time.Now()
	// return c.Status(201).JSON(data)

	Db := database.ConnectPostDB()
	Db.Delete(&data)
	return c.JSON(response.Response{Status: 201, Message: "Data successful inserted", Data: map[string]interface{}{"result": &data}})

}
