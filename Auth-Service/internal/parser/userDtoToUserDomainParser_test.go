package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/dtos"
	"reflect"
	"testing"
)

func TestNewUserDtoToUserDomainParser(t *testing.T) {
	tests := []struct {
		name string
		want IParser
	}{
		{
			name: "Test NewUserDtoToUserDomainParser",
			want: &userDtoToUserDomainParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDtoToUserDomainParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDtoToUserDomainParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDtoToUserDomainParser_Parser(t *testing.T) {
	type args struct {
		in []any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test userDto to userDomain success",
			args: args{
				in: []any{
					&dtos.User{
						Id:       "1",
						Name:     "User",
						Email:    "user@gmail.com",
						Password: "fujew9ru98re34",
					},
				},
			},
			want: &domain.User{
				Id:       "1",
				Name:     "User",
				Email:    "user@gmail.com",
				Password: "fujew9ru98re34",
			},
			wantErr: false,
		},
		{
			name: "Test userDto to userDomain fail",
			args: args{
				in: []any{"invalid input"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &userDtoToUserDomainParser{}
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
