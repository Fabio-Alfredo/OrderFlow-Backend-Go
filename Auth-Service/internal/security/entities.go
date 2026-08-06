package security

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserId           string `json:"user_id"`
	RegisteredClaims jwt.RegisteredClaims
}
