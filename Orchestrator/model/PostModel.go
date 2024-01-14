package model

import (
	"time"
)

type Post struct {
	Id          int
	Title       string
	Description string
	UserId      int
	CreatedAt   time.Time
	ModifiedAt  time.Time
	DeletedAt   time.Time
}
