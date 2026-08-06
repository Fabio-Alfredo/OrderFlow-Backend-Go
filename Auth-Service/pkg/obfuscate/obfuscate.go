package obfuscate

import (
	"Auth-Service/internal/domain/dtos"
	"Auth-Service/internal/domain/models"
	"strings"
)

const (
	charObfuscate     = "X"
	charObfuscatePass = "[REDACTED]"
)

func AuthCredentials(req dtos.LoginRequest) dtos.LoginRequest {
	req.Password = charObfuscatePass
	return req
}

func RegisterController(req dtos.RegisterRequest) dtos.RegisterRequest {
	req.User.Password = charObfuscatePass
	return req
}

func Password(req models.User) models.User {
	req.Password = charObfuscatePass
	return req
}

func TokenAuth(req models.Token) models.Token {
	req.Token = charObfuscate
	return req
}

// ObfuscateValue obfuscate the fields values
func ObfuscateValue(value, character string) string {
	return strings.Repeat(character, len(value))
}
