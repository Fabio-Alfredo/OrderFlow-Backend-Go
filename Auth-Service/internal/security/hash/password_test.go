package hash

import (
	"Auth-Service/internal/security"
	"Auth-Service/pkg/config"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_hashPassword_NewHashPassword(t *testing.T) {

	type fields struct {
		config config.IConfig
	}
	type args struct {
		config config.IConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   security.IHashMethods
	}{
		{
			name: "New HashPassword",
			fields: fields{
				config: nil,
			},
			want: &hashPassword{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashPassword{
				config: tt.fields.config,
			}
			if got := h.NewHashPassword(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashPassword_HashPassword(t *testing.T) {
	configs, _ := config.Load("../../../")

	type fields struct {
		config config.IConfig
	}
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Hash Password Success",
			fields: fields{
				config: configs,
			},
			args: args{
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashPassword{
				config: tt.fields.config,
			}
			got, err := h.Hash(tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == "" {
				t.Errorf("HashPassword() returned empty hash")
			}

			if got == tt.args.password {
				t.Errorf("HashPassword() returned plain password")
			}

			if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.password)); err != nil {
				t.Errorf("HashPassword() returned invalid bcrypt hash")
			}
		})
	}
}

func Test_hashPassword_ComparePassword(t *testing.T) {
	configs, _ := config.Load("../../../")

	type fields struct {
		config config.IConfig
	}
	type args struct {
		hash     string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Compare Password Success",
			fields: fields{
				config: configs,
			},
			args: args{
				hash:     "$2a$14$4KGQdJ9hDBh2hnzXjx3DHeZzCV/bOrRAEYQpXRUsdmjSAoY4Nellm",
				password: "123456",
			},
			want: true,
		},
		{
			name: "Compare Password Failed",
			fields: fields{
				config: configs,
			},
			args: args{
				hash:     "invalidhash",
				password: "123456",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashPassword{
				config: tt.fields.config,
			}
			if got := h.Compare(tt.args.hash, tt.args.password); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
