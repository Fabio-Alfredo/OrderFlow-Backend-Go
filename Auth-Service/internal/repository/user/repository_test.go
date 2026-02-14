package user

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	configs, _ := config.Load(".")
	log := logger.NewLogger()

	type args struct {
		config  config.IConfig
		sqlDb   *gorm.DB
		logger  logger.ILogger
		parsers parser.IFactory
	}
	tests := []struct {
		name string
		args args
		want repository.IUserRepository
	}{
		{
			name: "Test for New User Repository",
			args: args{
				config:  configs,
				sqlDb:   nil,
				logger:  log,
				parsers: nil,
			},
			want: &userRepository{
				config:  configs,
				db:      nil,
				logger:  log,
				parsers: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.config, tt.args.sqlDb, tt.args.logger, tt.args.parsers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
