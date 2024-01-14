package model

import (
	"time"
)

type Notifier struct {
	Id         int
	code       string
	UserId     int
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}
