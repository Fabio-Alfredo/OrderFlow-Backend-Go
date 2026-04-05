package token

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/repository/mocks"
	"Auth-Service/pkg/logger"
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestNewTokenRepository(t *testing.T) {
	log := logger.NewLogger()

	type args struct {
		sqlDb   *gorm.DB
		logger  logger.ILogger
		parsers parser.IFactory
	}
	tests := []struct {
		name string
		args args
		want repository.ITokenRepository
	}{
		{
			name: "Test for New Token Repository",
			args: args{
				sqlDb:   nil,
				logger:  log,
				parsers: nil,
			},
			want: &tokenRepository{
				db:      nil,
				logger:  log,
				parsers: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenRepository(tt.args.sqlDb, tt.args.logger, tt.args.parsers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenRepository_Save(t *testing.T) {
	log := logger.NewLogger()
	ctx := context.Background()

	//Mock for success save
	mockSucc, gdbSucc := mocks.GetDb()
	mockSucc.ExpectBegin()
	mockSucc.ExpectExec("").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mockSucc.ExpectCommit()

	//Mock error save
	mockErr, gdbErr := mocks.GetDb()
	mockErr.ExpectBegin()
	mockErr.ExpectExec("").
		WillReturnError(sqlmock.ErrCancelled)
	mockErr.ExpectRollback()

	parsers := factory.NewParserFactory()
	_ = parsers.Set(parser.TokenDomainToTokenRepositoryParser, parser.NewTokenDomainToTokenRepositoryParser())
	_ = parsers.Set(parser.TokenRepositoryToTokenDomainParser, parser.NewTokenRepositoryToTokenDomainParser())
	type fields struct {
		db      *gorm.DB
		logger  logger.ILogger
		parsers parser.IFactory
	}
	type args struct {
		ctx  context.Context
		data *domain.Token
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Token
		wantErr bool
	}{
		{
			name: "Test Save Token",
			fields: fields{
				db:      gdbSucc,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx: ctx,
				data: &domain.Token{
					Id:       "1",
					UserId:   "1",
					Token:    "token",
					IsActive: true,
				},
			},
			want: &domain.Token{
				Id:       "1",
				UserId:   "1",
				Token:    "token",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "Test Save Token Error",
			fields: fields{
				db:      gdbErr,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx: ctx,
				data: &domain.Token{
					Id:       "1",
					UserId:   "1",
					Token:    "token",
					IsActive: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &tokenRepository{
				db:      tt.fields.db,
				logger:  tt.fields.logger,
				parsers: tt.fields.parsers,
			}
			err := r.Save(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_tokenRepository_FindByUserAndActive(t *testing.T) {
	log := logger.NewLogger()
	ctx := context.Background()

	//Mock for success find
	mockSucc, gdbSucc := mocks.GetDb()
	rows := sqlmock.NewRows([]string{"id", "user_id", "token", "is_active"}).
		AddRow("1", "1", "token1", true).
		AddRow("2", "1", "token2", true)
	mockSucc.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `tokens` WHERE user_id = ? AND is_active = ? AND token = ?",
	)).
		WithArgs("1", true, "token1", 1).
		WillReturnRows(rows)

	//Mock error find
	mockErr, gdbErr := mocks.GetDb()
	mockErr.ExpectQuery("SELECT * FROM `tokens` WHERE user_id = ? AND is_active = ? AND token = ? ").
		WithArgs("1", true, "token1", 1).
		WillReturnError(sqlmock.ErrCancelled)

	parsers := factory.NewParserFactory()
	_ = parsers.Set(parser.TokenRepositoryToTokenDomainParser, parser.NewTokenRepositoryToTokenDomainParser())
	type fields struct {
		db      *gorm.DB
		logger  logger.ILogger
		parsers parser.IFactory
	}
	type args struct {
		ctx         context.Context
		userId      string
		active      bool
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Token
		wantErr bool
	}{
		{
			name: "Test FindByUserAndActive",
			fields: fields{
				db:      gdbSucc,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx:         ctx,
				userId:      "1",
				active:      true,
				tokenString: "token1",
			},
			want: &domain.Token{
				Id:       "1",
				UserId:   "1",
				Token:    "token1",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "Test FindByUserAndActive Error",
			fields: fields{
				db:      gdbErr,
				logger:  log,
				parsers: parsers,
			},
			args: args{
				ctx:         ctx,
				userId:      "1",
				active:      true,
				tokenString: "token1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &tokenRepository{
				db:      tt.fields.db,
				logger:  tt.fields.logger,
				parsers: tt.fields.parsers,
			}
			got, err := r.FindByUserAndActive(tt.args.ctx, tt.args.userId, tt.args.active, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUserAndActive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUserAndActive() got = %v, want %v", got, tt.want)
			}
		})
	}
}
