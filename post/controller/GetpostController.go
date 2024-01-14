package controller

import (
	"fmt"
	"post/database"
	"post/models"
	"post/response"

	"github.com/gofiber/fiber/v2"
)

func GetallPostController(c *fiber.Ctx) error {

	// define
	fmt.Println("Inside Get all Post controller")
	var data []models.Post

	Db := database.ConnectDb()
	query := Db.Raw("SELECT * FROM Posts")
	// if err := Db.Find(&data).Error; err != nil {
	// 	panic("Failed getting post data due to " + err.Error())
	// }
	if query.Scan(&data); query.Error != nil {
		panic(query.Error)
	}
	fmt.Println(data)
	return c.JSON(response.Response{Status: 200, Message: "Recevied all data", Data: map[string]interface{}{"data": data}})
}
