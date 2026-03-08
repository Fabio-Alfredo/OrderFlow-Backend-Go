package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"reflect"
	"testing"
)

func TestNewJWTService(t *testing.T) {
	log := logger.NewLogger()
	type args struct {
		config config.IConfig
		log    logger.ILogger
	}
	tests := []struct {
		name string
		args args
		want service.JWTMethods
	}{
		{
			name: "TestNewJWTService",
			args: args{
				config: nil,
				log:    log,
			},
			want: &jWTService{
				config: nil,
				log:    log,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWTService(tt.args.config, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jWTService_GenerateJWT(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()

	type fields struct {
		config config.IConfig
		log    logger.ILogger
	}
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "GenerateJWT Success",
			fields: fields{
				config: configs,
				log:    log,
			},
			args: args{
				user: &domain.User{
					Id: "1234",
				},
			},
			wantErr: false,
		},
		{
			name: "GenerateJWT Fail",
			fields: fields{
				config: configs,
				log:    log,
			},
			args: args{
				user: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &jWTService{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			_, err := s.GenerateJWT(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_jWTService_ValidateJWT1(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()

	ser := NewJWTService(configs, log)
	user := &domain.User{
		Id: "1234",
	}
	tokenString, _ := ser.GenerateJWT(user)

	type fields struct {
		config config.IConfig
		log    logger.ILogger
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ValidateJwt true",
			fields: fields{
				config: configs,
				log:    log,
			},
			args: args{
				tokenString,
			},
			want: true,
		},
		{
			name: "ValidateJwt false",
			fields: fields{
				config: configs,
				log:    log,
			},
			args: args{
				tokenString: "Dummy",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &jWTService{
				config: tt.fields.config,
				log:    tt.fields.log,
			}
			if got := s.ValidateJWT(tt.args.tokenString); got != tt.want {
				t.Errorf("ValidateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
