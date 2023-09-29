package main

import (
	"comments/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// init.Connect()
	db, err := gorm.Open(sqlite.Open("comment.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	db.AutoMigrate(&models.Comments{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world ! This is comment")
	})

	app.Listen(":3001")
}
