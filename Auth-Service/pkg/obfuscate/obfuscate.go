package obfuscate

import (
	"Auth-Service/internal/domain"
	"strings"
)

const (
	charObfuscate     = "X"
	charObfuscatePass = "[REDACTED]"
)

// Register Method for obfuscate password in register user
func Register(req domain.RegisterRequest) domain.RegisterRequest {
	req.User.Password = charObfuscatePass
	return req
}

// ObfuscateValue obfuscate the fields values
func ObfuscateValue(value, character string) string {
	return strings.Repeat(character, len(value))
}
