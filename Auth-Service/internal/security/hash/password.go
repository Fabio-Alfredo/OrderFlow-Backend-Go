package hash

import (
	"Auth-Service/internal/security"
	"Auth-Service/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

type hashPassword struct {
	config config.IConfig
}

func (h *hashPassword) NewHashPassword(config config.IConfig) security.IHashMethods {
	return &hashPassword{
		config: config,
	}
}

func (h *hashPassword) Hash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), h.config.GetInt("auth.secure.hash_cost"))
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func (h *hashPassword) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
