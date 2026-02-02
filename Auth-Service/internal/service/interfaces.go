package service

import (
	"Auth-Service/internal/dtos"
	"context"
)

type IAuthService interface {
	Register(ctx context.Context, user *dtos.User) *RegisterServiceResp
	Login(ctx context.Context, req *dtos.LoginRequest) *dtos.LoginResponse
}

type IParser interface {
	Parser(in ...any) (any, error)
}

type IFactory interface {
	Set(key string, parser IParser) error
	Get(key string) (parser IParser, err error)
}
