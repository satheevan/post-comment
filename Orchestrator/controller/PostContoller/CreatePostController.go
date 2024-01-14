package postcontroller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/controller"
	"orchestrtorservices.com/orchestrator/response"
)

type PostController struct {
	controller.BaseController
}

func (pc PostController) CreatePostController(c *fiber.Ctx) error {
	// return c.Status(201).JSON(data)

	// Db := database.ConnectPostDB()
	// Db.Create(&data)
	// return c.JSON(response.Response{Status: 201, Message: "Data successful inserted", Data: map[string]interface{}{"result": &data}})

	// microservices
	// Make an API Get Request to another service
	res, err := http.Post("http://localhost:3003/api/createpost", "application/json", bytes.NewBuffer(c.Body()))
	if err != nil {
		fmt.Println("API Request Faild", err)
		return c.Status(http.StatusInternalServerError).SendString("Faild to call the Post services")
	}
	defer res.Body.Close()
	//Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("faild to read Response:", err)
		return c.Status(http.StatusInternalServerError).SendString("Faild to read the response from other service")
	}
	fmt.Println("Response Body :", string(body))
	return c.JSON(response.Response{Status: 201, Message: "Data successful inserted", Data: map[string]interface{}{"result": string(body)}})
}
