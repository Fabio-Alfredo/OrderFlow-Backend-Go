package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"reflect"
	"testing"
)

func TestNewUserRepositoryToUserDomainParser(t *testing.T) {
	tests := []struct {
		name string
		want IParser
	}{
		{
			name: "Test NewUserRepositoryToUserDomainParser",
			want: &userRepositoryToUserDomainParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepositoryToUserDomainParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepositoryToUserDomainParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepositoryToUserDomainParser_Parser(t *testing.T) {
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
			name: "Test userRepository to userDomain success",
			args: args{
				in: []any{&repository.User{
					Name:     "User",
					Email:    "user@gmai.com",
					Password: "fujew9ru98re34",
					Status:   "activo",
				}},
			},
			want: &domain.User{
				Name:     "User",
				Email:    "user@gmai.com",
				Password: "fujew9ru98re34",
				Status:   "activo",
			},
			wantErr: false,
		},
		{
			name: "Test userRepository to userDomain invalid input",
			args: args{
				in: []any{"invalid input"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &userRepositoryToUserDomainParser{}
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
