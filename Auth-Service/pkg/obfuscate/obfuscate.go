package obfuscate

import (
	"Auth-Service/internal/dtos"
	"strings"
)

const (
	charObfuscate     = "X"
	charObfuscatePass = "[REDACTED]"
)

// Register Method for obfuscate password in register user
func Register(user dtos.User) dtos.User {
	user.Password = charObfuscatePass
	return user
}

// ObfuscateValue obfuscate the fields values
func ObfuscateValue(value, character string) string {
	return strings.Repeat(character, len(value))
}
