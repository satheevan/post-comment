package controller

import (
	"fmt"
	"post/database"
	"post/models"
	"post/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePostController(c *fiber.Ctx) error {
	fmt.Println("inside CreatePostController")
	data := new(models.Post)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data.CreatedAt = time.Now()
	fmt.Println(data)
	Db := database.ConnectDb()
	Db.Create(&data)
	return c.JSON(response.Response{Status: 201, Message: "Data Inserted successfully", Data: map[string]interface{}{"data": &data}})
}
