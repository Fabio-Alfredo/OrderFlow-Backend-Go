package obfuscate

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/dtos"
	"strings"
)

const (
	charObfuscate     = "X"
	charObfuscatePass = "[REDACTED]"
)

// Register Method for obfuscate password in register user
func RegisterService(user domain.User) domain.User {
	user.Password = charObfuscatePass
	return user
}

func RegisterController(req dtos.RegisterRequest) dtos.RegisterRequest {
	req.User.Password = charObfuscatePass
	return req
}

// ObfuscateValue obfuscate the fields values
func ObfuscateValue(value, character string) string {
	return strings.Repeat(character, len(value))
}
