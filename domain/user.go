package domain

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Password  string
	Age       int
	Height    int
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

type Users []User
