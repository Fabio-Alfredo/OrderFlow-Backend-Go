package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(in string, hashCost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in), hashCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(in, inHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inHash), []byte(in))
	return err == nil
}
