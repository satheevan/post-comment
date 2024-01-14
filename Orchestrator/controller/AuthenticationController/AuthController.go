package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/controller"
	"orchestrtorservices.com/orchestrator/database"
	"orchestrtorservices.com/orchestrator/model"
	"orchestrtorservices.com/orchestrator/response"
)

type AuthenticationUser struct {
	controller.BaseController
}

func (uc AuthenticationUser) Register(c *fiber.Ctx) error {
	var data = model.AuthenticationUser{}

	if err := c.BodyParser(&data); err != nil {
		return (err)
	}
	fmt.Println(&data)
	data.CreatedAt = time.Now()

	Db := database.ConnectUserDB()
	Db.Create(&data)
	return c.JSON(response.Response{Status: 201, Message: "Successfully Register", Data: map[string]interface{}{"result": &data}})

}
