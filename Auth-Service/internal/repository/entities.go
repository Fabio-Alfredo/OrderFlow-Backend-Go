package repository

import (
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Id       string    `gorm:"primary_key;column:id"`
	Name     string    `gorm:"column:username"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password_hash"`
	Status   string    `gorm:"column:status"`
	CreateAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string { return "users" }

type Token struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token_has"`
	ExpiresAt time.Time `json:"expires_at"`
	IsActive  bool      `json:"is_active"`
	TimesTamp time.Time `json:"times_tamp;autoUpdateTime"`
	CreateAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Token) TableName() string { return "tokens" }
