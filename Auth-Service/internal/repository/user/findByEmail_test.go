package user

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

func Test_userRepository_FindEmail(t *testing.T) {
	date, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	configs, _ := config.Load(".")
	log := logger.NewLogger()
	ctx := context.Background()
	errorDummy := errors.New("dummy error")

	//Mock for success find
	mockSucc, gdbSucc := mocks.GetDb()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "status", "create_at", "update_at"}).
		AddRow("1234", "user", "user@gmail.com", "kjf4490", "activo", date, date)
	mockSucc.ExpectQuery("").
		WithArgs("user@gmail.com", 1).
		WillReturnRows(rows)

	//Mock for record not found
	mockErr, gdbErr := mocks.GetDb()
	mockErr.ExpectQuery("").
		WithArgs("not-email", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	//Mock for repository error
	mockErrR, gdbErrR := mocks.GetDb()
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
			want: &repository.User{
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
