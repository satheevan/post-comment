package init

import (
	// folders
	"comments/models"

	//Package
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("comment.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	db.AutoMigrate(&models.Comments{})
}
