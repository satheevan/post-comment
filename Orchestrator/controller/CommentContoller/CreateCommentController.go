package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"orchestrtorservices.com/orchestrator/controller"
	"orchestrtorservices.com/orchestrator/response"
)

type CommentController struct {
	controller.BaseController
}

func (cc CommentController) CreateCommentController(c *fiber.Ctx) error {

	// //Parse the incoming  Json Request
	// var data model.Comment
	// fmt.Println(data)
	// if err := c.BodyParser(&data); err != nil {
	// 	return err
	// }
	// fmt.Println(data)
	// data.CreatedAt = time.Now()
	// //database
	// Db := database.ConnectCommentDB()
	// Db.Create(&data)

	// return c.JSON(response.Response{Status: 201, Message: "Successfully Inserted", Data: map[string]interface{}{"result": &data}})
	// -------------------------------------------------------------------------------------------
	// microservices
	//  Make an API request call to Another services
	res, err := http.Post("http://localhost:3004/api/createcomment", "application/json", bytes.NewBuffer(c.Body()))

	if err != nil {
		fmt.Println("API request faild", err)
		return c.Status(http.StatusInternalServerError).SendString("faild to call the comment services")
	}
	defer res.Body.Close()
	//read the rresponse body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Faild to read the response", err)
		return c.Status(http.StatusInternalServerError).SendString("faild to read the response serive from other serrives")
	}
	fmt.Println("Response Body :", string(body))
	return c.JSON(response.Response{Status: 201, Message: "Data successful inserted", Data: map[string]interface{}{"data": string(body)}})
}
