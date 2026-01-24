package impl

import (
	"Auth-Service/internal/repository"
	"Auth-Service/internal/repository/mocks"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	configs, _ := config.Load(".")
	log := logger.NewLogger()

	type args struct {
		config config.IConfig
		sqlDb  *gorm.DB
		logger logger.ILogger
	}
	tests := []struct {
		name string
		args args
		want repository.IUserRepository
	}{
		{
			name: "Test for New User Repository",
			args: args{
				config: configs,
				sqlDb:  nil,
				logger: log,
			},
			want: &userRepository{
				config: configs,
				db:     nil,
				logger: log,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.config, tt.args.sqlDb, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Save(t *testing.T) {
	configs, _ := config.Load(".")
	log := logger.NewLogger()
	ctx := context.Background()
	errorDummy := errors.New("dummy error")

	//Mock for error
	mockErr, gdbErr := mocks.GetDb(configs)
	mockErr.ExpectBegin()
	mockErr.ExpectExec("").
		WillReturnError(errorDummy)
	mockErr.ExpectRollback()

	//Mock for success
	mockSucc, gdbSucc := mocks.GetDb(configs)
	mockSucc.ExpectBegin()
	mockSucc.ExpectExec("").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockSucc.ExpectCommit()

	type fields struct {
		config config.IConfig
		db     *gorm.DB
		logger logger.ILogger
	}
	type args struct {
		ctx  context.Context
		data *repository.User
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
				config: configs,
				db:     gdbErr,
				logger: log,
			},
			args: args{
				ctx: ctx,
				data: &repository.User{
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
				config: configs,
				db:     gdbSucc,
				logger: log,
			},
			args: args{
				ctx: ctx,
				data: &repository.User{
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
				config: tt.fields.config,
				db:     tt.fields.db,
				logger: tt.fields.logger,
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

func Test_userRepository_FindEmail(t *testing.T) {
	date, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	configs, _ := config.Load(".")
	log := logger.NewLogger()
	ctx := context.Background()
	errorDummy := errors.New("dummy error")

	//Mock for success find
	mockSucc, gdbSucc := mocks.GetDb(configs)
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "status", "create_at", "update_at"}).
		AddRow("1234", "user", "user@gmail.com", "kjf4490", "activo", date, date)
	mockSucc.ExpectQuery("").
		WithArgs("user@gmail.com", 1).
		WillReturnRows(rows)

	//Mock for record not found
	mockErr, gdbErr := mocks.GetDb(configs)
	mockErr.ExpectQuery("").
		WithArgs("not-email", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	//Mock for repository error
	mockErrR, gdbErrR := mocks.GetDb(configs)
	mockErrR.ExpectQuery("").
		WithArgs("not-found", 1).
		WillReturnError(errorDummy)

	type fields struct {
		config config.IConfig
		db     *gorm.DB
		logger logger.ILogger
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test for find user Success",
			fields: fields{
				config: configs,
				db:     gdbSucc,
				logger: log,
			},
			args: args{
				ctx:   ctx,
				email: "user@gmail.com",
			},
			want: repository.User{
				Id:       "1234",
				Name:     "user",
				Email:    "user@gmail.com",
				Password: "kjf4490",
				Status:   "activo",
				CreateAt: date,
				UpdateAt: date,
			},
			wantErr: false,
		},
		{
			name: "Test for find user Error",
			fields: fields{
				config: configs,
				db:     gdbErr,
				logger: log,
			},
			args: args{
				ctx:   ctx,
				email: "not-email",
			},
			want:    repository.ErrUserNotFound,
			wantErr: true,
		},
		{
			name: "Test for find user Error repository",
			fields: fields{
				config: configs,
				db:     gdbErrR,
				logger: log,
			},
			args: args{
				ctx:   ctx,
				email: "not-found",
			},
			want:    errorDummy,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepository{
				config: tt.fields.config,
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.FindEmail(tt.args.ctx, tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !errors.Is(err, tt.want.(error)) {
					t.Errorf("FindEmail() error = %v, want %v", err, tt.want)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindEmail() got = %v, want %v", got, tt.want)
			}

		})
	}
}
