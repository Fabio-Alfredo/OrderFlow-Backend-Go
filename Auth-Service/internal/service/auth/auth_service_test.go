package auth

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"reflect"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		config     config.IConfig
		log        logger.ILogger
		repository repository.IUserRepository
		parsers    parser.IFactory
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
				parsers:    nil,
			},
			want: &authService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.config, tt.args.log, tt.args.repository, tt.args.parsers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}
