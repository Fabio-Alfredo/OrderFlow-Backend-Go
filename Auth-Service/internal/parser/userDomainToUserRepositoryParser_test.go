package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/config"
	"reflect"
	"testing"
)

func TestNewUserDtoToUserRepositoryParser(t *testing.T) {
	configs, _ := config.Load("../../../")
	type args struct {
		config config.IConfig
	}
	tests := []struct {
		name string
		args args
		want IParser
	}{
		{
			name: "Test NewUserDtoToUserRepositoryParser",
			args: args{
				config: configs,
			},
			want: &userDomainToUserRepositoryParser{
				config: configs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDomainToUserRepositoryParser(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDtoToUserRepositoryParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDtoToUserRepositoryParser_Parser(t *testing.T) {
	configs, _ := config.Load("../../")
	type fields struct {
		config config.IConfig
	}
	type args struct {
		in []any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test userDto to userRepository success",
			fields: fields{
				config: configs,
			},
			args: args{
				in: []any{
					&domain.User{
						Id:       "1",
						Name:     "user",
						Email:    "user@gmail.com",
						Password: "dfserrt4325",
					},
				},
			},
			want: &repository.User{
				Id:       "1",
				Name:     "user",
				Email:    "user@gmail.com",
				Password: "dfserrt4325",
				Status:   configs.GetString("auth.registration.default.status"),
			},
			wantErr: false,
		},
		{
			name:   "Test userDto to userRepository invalid input",
			fields: fields{},
			args: args{
				in: []any{
					"invalid input",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &userDomainToUserRepositoryParser{
				config: tt.fields.config,
			}
			got, err := p.Parser(tt.args.in...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
