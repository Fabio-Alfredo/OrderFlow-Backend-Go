package user

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/internal/repository/mocks"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func Test_userRepository_Save(t *testing.T) {
	configs, _ := config.Load("../../../")
	log := logger.NewLogger()
	ctx := context.Background()
	errorDummy := errors.New("dummy error")

	//Mock for error
	mockErr, gdbErr := mocks.GetDb()
	mockErr.ExpectBegin()
	mockErr.ExpectExec("").
		WillReturnError(errorDummy)
	mockErr.ExpectRollback()

	//Mock for success
	mockSucc, gdbSucc := mocks.GetDb()
	mockSucc.ExpectBegin()
	mockSucc.ExpectExec("").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockSucc.ExpectCommit()

	parsers := factory.NewParserFactory()
	userRepoParser := parser.NewUserDomainToUserRepositoryParser(configs)
	_ = parsers.Set(parser.UserDomainToUserRepositoryParser, userRepoParser)

	type fields struct {
		config  config.IConfig
		db      *gorm.DB
		logger  logger.ILogger
		parsers parser.IFactory
	}
	type args struct {
		ctx  context.Context
		data *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "Test for Save User Error",
			fields: fields{
				config:  configs,
				db:      gdbErr,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx: ctx,
				data: &domain.User{
					Name:     "User",
					Email:    "user@gmail.com",
					Password: "fujew9ru98re34",
					Status:   "activo",
				},
			},
			want:    errorDummy,
			wantErr: true,
		},
		{
			name: "Test for Save User Success",
			fields: fields{
				config:  configs,
				db:      gdbSucc,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx: ctx,
				data: &domain.User{
					Name:     "User",
					Email:    "user@gmail.com",
					Password: "fujew9ru98re34",
					Status:   "activo",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepository{
				config:  tt.fields.config,
				db:      tt.fields.db,
				logger:  tt.fields.logger,
				parsers: tt.fields.parsers,
			}
			got := r.Save(tt.args.ctx, tt.args.data)
			if (got != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", got, tt.wantErr)
			}

			if tt.wantErr && tt.want != nil {
				if !errors.Is(got, tt.want) {
					t.Errorf("Save() error = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
