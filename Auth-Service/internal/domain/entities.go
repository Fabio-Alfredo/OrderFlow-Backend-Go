package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
	Status   string
}

type AuthCredentials struct {
	Identifier string
	Password   string
}

type RegisterResult struct {
	Code    string
	Message string
}

type LoginResult struct {
	Token       string
	Description string
}

type Token struct {
	Id        string
	UserId    string
	Token     string
	ExpiresAt time.Time
	IsActive  bool
	TimesTamp time.Time
}

type JWTClaims struct {
	UserId string
	jwt.RegisteredClaims
}
