package database

import (
	"Auth-Service/internal/database/mocks"
	"Auth-Service/pkg/config"
	"database/sql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewSQLConfig(t *testing.T) {
	configuration, _ := config.Load("../../")
	type args struct {
		config config.IConfig
	}
	tests := []struct {
		name string
		args args
		want ISqlDB
	}{
		{
			name: "Test NewSQLConfig",
			args: args{
				config: configuration,
			},
			want: &sqlConfig{
				config: configuration,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlConfig_GetDB(t *testing.T) {
	configs, _ := config.Load("../../")
	type fields struct {
		config config.IConfig
		db     *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "Test connection to database existing",
			fields: fields{
				config: configs,
				db:     &gorm.DB{},
			},
			want:    &gorm.DB{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlConfig{
				config: tt.fields.config,
				db:     tt.fields.db,
			}
			got, err := s.GetDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlConfig_openConnection(t *testing.T) {
	configs, _ := config.Load("../../")

	type fields struct {
		config config.IConfig
		db     *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sql.DB
		wantErr bool
	}{
		{
			name: "Test connection Success",
			fields: fields{
				config: configs,
				db:     nil,
			},
			want:    mocks.GetDbConnectionMock(configs),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlConfig{
				config: tt.fields.config,
				db:     tt.fields.db,
			}
			_, err := s.openConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("openConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_sqlConfig_getDns(t *testing.T) {
	configs, _ := config.Load("../../")
	type fields struct {
		config config.IConfig
		db     *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test getDns",
			fields: fields{
				config: configs,
				db:     nil,
			},
			want: mocks.GetDnsMock(configs),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlConfig{
				config: tt.fields.config,
				db:     tt.fields.db,
			}
			if got := s.getDns(); got != tt.want {
				t.Errorf("getDns() = %v, want %v", got, tt.want)
			}
		})
	}
}
