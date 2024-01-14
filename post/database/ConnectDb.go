package database

import (
	"post/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDb() gorm.DB {
	Db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	Db.AutoMigrate(&models.Post{})
	return *Db
}
