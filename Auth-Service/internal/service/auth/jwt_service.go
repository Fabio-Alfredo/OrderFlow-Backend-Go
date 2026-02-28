package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtServiceMethodsTitle = "jwtService: "
)

type jWTService struct {
	config config.IConfig
	log    logger.ILogger
}

func NewJWTService(config config.IConfig, log logger.ILogger) service.JWTMethods {
	return &jWTService{
		config: config,
		log:    log,
	}
}

func (s *jWTService) GenerateJWT(user *domain.User) (string, error) {

	claims, err := s.buildClaims(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.GetString("auth.jwt.secret")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *jWTService) ValidateJWT(tokenString string) (*domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.GetString("auth.jwt.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.JWTClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrNotFound
	}

	return claims, nil
}

func (s *jWTService) buildClaims(user *domain.User) (*domain.JWTClaims, error) {

	if user == nil {
		return nil, domain.ErrNotFound
	}

	duration, err := time.ParseDuration(s.config.GetString("auth.jwt.expiration"))
	if err != nil {
		return nil, err
	}

	claims := &domain.JWTClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.config.GetString("auth.jwt.issuer"),
		},
	}

	return claims, nil
}
