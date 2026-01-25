package security

import (
	"Auth-Service/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
	config config.IConfig
}

func NewHasher(config config.IConfig) IHash {
	return &Hasher{
		config: config,
	}
}

func (h *Hasher) HashPassword(in string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in), h.config.GetInt("auth.secure.hash_cost"))

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *Hasher) CheckPasswordHash(in, inHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inHash), []byte(in))
	return err == nil
}
