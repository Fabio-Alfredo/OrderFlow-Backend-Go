package service

import "Auth-Service/internal/domain"

type IAuthService interface {
	Register(req *domain.RegisterRequest) *domain.RegisterResponse
	Login(req *domain.LoginRequest) *domain.LoginResponse
}
