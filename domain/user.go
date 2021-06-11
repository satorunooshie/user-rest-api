package domain

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Age       int       `json:"age"`
	Height    int       `json:"height"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type Users []User
