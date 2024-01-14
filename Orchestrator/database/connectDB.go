package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"orchestrtorservices.com/orchestrator/model"
)

func ConnectPostDB() gorm.DB {

	Db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	Db.AutoMigrate(&model.Post{})
	return *Db
}

func ConnectCommentDB() *gorm.DB {

	Db, err := gorm.Open(sqlite.Open("comment.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	Db.AutoMigrate(&model.Comment{})
	return Db
}

func ConnectUserDB() *gorm.DB {

	Db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})

	if err != nil {
		panic("faild to connect database")
	}
	Db.AutoMigrate(&model.AuthenticationUser{})
	return Db
}
