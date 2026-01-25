package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"reflect"
	"testing"
)

func TestNewUserDomainToUserPersistence(t *testing.T) {
	tests := []struct {
		name string
		want service.IParser
	}{
		{
			name: "Test new Parser UserDomainToUserPersistence",
			want: &userDomainToUserPersistence{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDomainToUserPersistence(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDomainToUserPersistence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDomainToUserPersistence_Parser(t *testing.T) {
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
			name: "Test for Success parser userDomainToUserPersistence",
			args: args{
				in: []any{
					&domain.User{
						Id:       "550e8400-e29b-41d4-a716-446655440000",
						Name:     "user",
						Email:    "user@gmail.com",
						Password: "dfsadsfjo34j48934",
					},
				},
			},
			want: &repository.User{
				Id:       "550e8400-e29b-41d4-a716-446655440000",
				Name:     "user",
				Email:    "user@gmail.com",
				Password: "dfsadsfjo34j48934",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &userDomainToUserPersistence{}
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
