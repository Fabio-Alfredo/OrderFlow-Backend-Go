package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"reflect"
	"testing"
)

func TestNewTokenDomainToTokenRepositoryParser(t *testing.T) {
	tests := []struct {
		name string
		want IParser
	}{
		{
			name: "Test NewTokenDomainToTokenRepositoryParser",
			want: &tokenDomainToTokenRepositoryParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenDomainToTokenRepositoryParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenDomainToTokenRepositoryParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenDomainToTokenRepositoryParser_Parser(t *testing.T) {
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
			name: "Test tokenDomain to tokenRepository success",
			args: args{
				in: []any{&domain.Token{
					Id:       "1",
					UserId:   "1",
					Token:    "fujew9ru98re34",
					IsActive: true,
				}},
			},
			want: &repository.Token{
				Id:       "1",
				UserId:   "1",
				Token:    "fujew9ru98re34",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "Test tokenDomain to tokenRepository invalid input",
			args: args{
				in: []any{"invalid input"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tokenDomainToTokenRepositoryParser{}
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
