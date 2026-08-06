package token

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/security"
	"Auth-Service/internal/service"
	"Auth-Service/internal/service/token/mocks"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"context"
	"reflect"
	"testing"
)

func TestNewTokenService(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()

	type args struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.ITokenRepository
		jwtMethods security.IJWTMethods
	}
	tests := []struct {
		name string
		args args
		want service.ITokenService
	}{
		{
			name: "Test New TokenService",
			args: args{
				config: configs,
				log:    log,
			},
			want: &tokenService{
				config: configs,
				log:    log,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenService(tt.args.config, tt.args.log, tt.args.repository, tt.args.jwtMethods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenService_Register(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.ITokenRepository
		jwtMethods security.IJWTMethods
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Register Token Success",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewTokenRepositoryMock(false),
				jwtMethods: mocks.NewJwtMethodsMock(false),
			},
			args: args{
				ctx: ctx,
				user: &models.User{
					Id:       "1",
					Name:     "Test",
					Email:    "test@example.com",
					Password: "123456",
					Status:   "active",
				},
			},
			want:    "token",
			wantErr: false,
		},
		{
			name: "Test Register Token JWT Error",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewTokenRepositoryMock(false),
				jwtMethods: mocks.NewJwtMethodsMock(true),
			},
			args: args{
				ctx: ctx,
				user: &models.User{
					Id:       "1",
					Name:     "Test",
					Email:    "test@example.com",
					Password: "123456",
					Status:   "active",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tokenService{
				config:     tt.fields.config,
				log:        tt.fields.log,
				repository: tt.fields.repository,
				jwtMethods: tt.fields.jwtMethods,
			}
			got, err := s.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenService_IsValid(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()

	type fields struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.ITokenRepository
		jwtMethods security.IJWTMethods
	}
	type args struct {
		ctx         context.Context
		tokenString string
		userId      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test Valid Token",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewTokenRepositoryMock(false),
				jwtMethods: mocks.NewJwtMethodsMock(false),
			},
			args: args{
				ctx:         ctx,
				tokenString: "token",
				userId:      "1",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Invalid Token",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewTokenRepositoryMock(false),
				jwtMethods: mocks.NewJwtMethodsMock(true),
			},
			args: args{
				ctx:         ctx,
				tokenString: "token",
				userId:      "2",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test Repository Error",
			fields: fields{
				config:     configs,
				log:        log,
				repository: mocks.NewTokenRepositoryMock(true),
				jwtMethods: mocks.NewJwtMethodsMock(false),
			},
			args: args{
				ctx:         ctx,
				tokenString: "token",
				userId:      "1",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tokenService{
				config:     tt.fields.config,
				log:        tt.fields.log,
				repository: tt.fields.repository,
				jwtMethods: tt.fields.jwtMethods,
			}
			got, err := s.IsValid(tt.args.ctx, tt.args.tokenString, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
