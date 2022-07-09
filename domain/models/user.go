package models

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-" column:"password"`
	Gender    string    `json:"gender"`
	Phone     int64     `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" column:"password"`
	Gender   string `json:"gender"`
	Phone    int64  `json:"phone"`
	Address  string `json:"address"`
}

type UserResponse struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-" column:"password"`
	Gender    string    `json:"gender"`
	Phone     int64     `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
