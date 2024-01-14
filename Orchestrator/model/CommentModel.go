package model

import (
	"time"
)

type Comment struct {
	Id         int
	Text       string
	PostId     int
	UserId     int   
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}
