package security

import "Auth-Service/internal/domain/models"

type IJWTMethods interface {
	GenerateJWT(user *models.User) (string, error)
	ValidateJWT(token string) bool
	GetClaims(tokenString string) (*JWTClaims, error)
}

type IHashMethods interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}
