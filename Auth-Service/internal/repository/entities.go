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
	Id        string    `gorm:"primary_key;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Token     string    `gorm:"column:token"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	IsActive  bool      `gorm:"column:is_active"`
	TimesTamp time.Time `gorm:"column:times_tamp;autoUpdateTime"`
	CreateAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Token) TableName() string { return "tokens" }
