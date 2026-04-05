package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"fmt"
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

func NewJWTService(config config.IConfig, log logger.ILogger) service.IJWTMethods {
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

func (s *jWTService) ValidateJWT(tokenString string) bool {
	_, err := s.parseToken(tokenString, &jwt.MapClaims{})

	if err != nil {
		return false
	}

	return true
}

func (s *jWTService) GetClaims(tokenString string) (*domain.JWTClaims, error) {
	claims := &domain.JWTClaims{}

	_, err := s.parseToken(tokenString, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *jWTService) parseToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(s.config.GetString("auth.jwt.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
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
