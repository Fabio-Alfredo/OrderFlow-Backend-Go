package console

import (
	"Auth-Service/internal/dtos"
	"context"
)

const (
	ContextKeyRegisterId = "email"
)

func SetContextWithRegister(ctx context.Context, request *dtos.RegisterRequest) context.Context {
	ctx = context.WithValue(ctx, ContextKeyRegisterId, request.User.Email)

	return ctx
}
