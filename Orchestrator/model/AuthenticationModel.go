package model

import "time"

type AuthenticationUser struct {
	Id         int
	UserName   string
	MobileNo   int
	Email      string
	NotifierId int
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}
