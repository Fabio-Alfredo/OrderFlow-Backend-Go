package auth

import (
	"Auth-Service/internal/domain/dtos"
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/security"
	"Auth-Service/internal/service"
	mocks2 "Auth-Service/internal/service/auth/mocks"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/uuid"
	"context"
	"reflect"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		service    service.ITokenService
		uuid       uuid.IGenerator
		jwt        security.IJWTMethods
		password   security.IHashMethods
	}
	tests := []struct {
		name string
		args args
		want service.IAuthService
	}{
		{
			name: "Test NewAuthService",
			args: args{
				config:     nil,
				log:        nil,
				repository: nil,
				service:    nil,
				uuid:       nil,
				jwt:        nil,
				password:   nil,
			},
			want: &authService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.config, tt.args.log, tt.args.repository, tt.args.service, tt.args.uuid, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_Register(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		service    service.ITokenService
		uuid       uuid.IGenerator
		password   security.IHashMethods
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test register user success",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks2.NewUserRepositoryMock(false, false),
				service:    nil,
				uuid:       mocks2.NewUUIDMock(),
				password:   mocks2.NewHashMock(false),
			},
			args: args{
				ctx: ctx,
				user: &models.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			wantErr: false,
		},
		{
			name: "Test register user error internal",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks2.NewUserRepositoryMock(true, false),
				service:    nil,
				uuid:       mocks2.NewUUIDMock(),
				password:   mocks2.NewHashMock(false),
			},
			args: args{
				ctx: ctx,
				user: &models.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			wantErr: true,
		},
		{
			name: "Test register user error already exist user by email",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks2.NewUserRepositoryMock(false, true),
				service:    nil,
				uuid:       mocks2.NewUUIDMock(),
				password:   mocks2.NewHashMock(false),
			},
			args: args{
				ctx: ctx,
				user: &models.User{
					Id:       "",
					Name:     "user",
					Email:    "user@gmail.com",
					Password: "aser34f34qf",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				config:     tt.fields.config,
				log:        tt.fields.log,
				repository: tt.fields.repository,
				service:    tt.fields.service,
				uuid:       tt.fields.uuid,
				password:   tt.fields.password,
			}
			err := s.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_authService_Login(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		service    service.ITokenService
		uuid       uuid.IGenerator
		password   security.IHashMethods
	}
	type args struct {
		ctx             context.Context
		authCredentials *dtos.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test login user success",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks2.NewUserRepositoryMock(false, true),
				service:    mocks2.NewTokenServiceMock(false),
				uuid:       nil,
				password:   mocks2.NewHashMock(false),
			},
			args: args{
				ctx: ctx,
				authCredentials: &dtos.LoginRequest{
					Identifier: "user",
					Password:   "aser34f34qf",
				},
			},
			want:    "token",
			wantErr: false,
		},
		{
			name: "Test login user error internal",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks2.NewUserRepositoryMock(true, true),
				service:    mocks2.NewTokenServiceMock(false),
				uuid:       nil,
				password:   mocks2.NewHashMock(true),
			},
			args: args{
				ctx: ctx,
				authCredentials: &dtos.LoginRequest{
					Identifier: "user",
					Password:   "aser34f34qf",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				config:     tt.fields.config,
				log:        tt.fields.log,
				repository: tt.fields.repository,
				service:    tt.fields.service,
				uuid:       tt.fields.uuid,
				password:   tt.fields.password,
			}
			got, err := s.Login(tt.args.ctx, tt.args.authCredentials)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
