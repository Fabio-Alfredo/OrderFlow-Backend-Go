package auth

import (
	"Auth-Service/internal/dtos"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/internal/service/mocks"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"context"
	"reflect"
	"testing"
)

func Test_authService_Register(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()

	parsers := factory.NewParserFactory()
	_ = parsers.Set(parser.UserDtoToUserRepositoryParser, parser.NewUserDtoToUserRepositoryParser(configs))

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		parsers    service.IFactory
	}
	type args struct {
		ctx     context.Context
		userDto *dtos.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *service.RegisterServiceResp
	}{
		{
			name: "Test register user success",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(false, false),
				parsers:    parsers,
			},
			args: args{
				ctx: ctx,
				userDto: &dtos.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want: &service.RegisterServiceResp{
				Code:    configs.GetString("auth.register.success.code"),
				Message: configs.GetString("auth.register.success.message"),
			},
		},
		{
			name: "Test register user error internal",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(true, false),
				parsers:    parsers,
			},
			args: args{
				ctx: ctx,
				userDto: &dtos.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want: &service.RegisterServiceResp{
				Code:    configs.GetString("auth.register.errors.INTERNAL.code"),
				Message: "error",
			},
		},
		{
			name: "Test register user error already exist user by email",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(false, true),
				parsers:    parsers,
			},
			args: args{
				ctx: ctx,
				userDto: &dtos.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want: &service.RegisterServiceResp{
				Code:    configs.GetString("auth.register.errors.USER_ALREADY_EXISTS.code"),
				Message: configs.GetString("auth.register.errors.USER_ALREADY_EXISTS.message"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				config:     tt.fields.config,
				log:        tt.fields.log,
				repository: tt.fields.repository,
				parsers:    tt.fields.parsers,
			}
			if got := s.Register(tt.args.ctx, tt.args.userDto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
