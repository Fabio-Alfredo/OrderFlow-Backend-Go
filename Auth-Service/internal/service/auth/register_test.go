package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
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

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		parsers    parser.IFactory
	}
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.RegisterResult
		wantErr bool
	}{
		{
			name: "Test register user success",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(false, false),
				parsers:    nil,
			},
			args: args{
				ctx: ctx,
				user: &domain.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want: &domain.RegisterResult{
				Code:    configs.GetString("auth.register.success.code"),
				Message: configs.GetString("auth.register.success.message"),
			},
			wantErr: false,
		},
		{
			name: "Test register user error internal",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(true, false),
				parsers:    nil,
			},
			args: args{
				ctx: ctx,
				user: &domain.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test register user error already exist user by email",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewMockRepository(false, true),
				parsers:    nil,
			},
			args: args{
				ctx: ctx,
				user: &domain.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			want:    nil,
			wantErr: true,
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
			got, err := s.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
