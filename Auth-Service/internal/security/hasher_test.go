package security

import (
	"Auth-Service/pkg/config"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewHasher(t *testing.T) {
	configs, _ := config.Load("../../")

	type args struct {
		config config.IConfig
	}
	tests := []struct {
		name string
		args args
		want IHash
	}{
		{
			name: "Test for NewHash",
			args: args{
				config: configs,
			},
			want: &Hasher{config: configs},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHasher(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHasher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_HashPassword(t *testing.T) {
	configs, _ := config.Load("../../")

	type fields struct {
		config config.IConfig
	}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "HashPassword success",
			fields: fields{config: configs},
			args: args{
				in: "123456",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{
				config: tt.fields.config,
			}

			got, err := h.HashPassword(tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == "" {
				t.Errorf("HashPassword() returned empty hash")
			}

			if got == tt.args.in {
				t.Errorf("HashPassword() returned plain password")
			}

			if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.in)); err != nil {
				t.Errorf("HashPassword() returned invalid bcrypt hash")
			}
		})
	}
}

func TestHasher_CheckPasswordHash(t *testing.T) {
	configs, _ := config.Load("../../")

	hasher := &Hasher{
		config: configs,
	}

	validPassword := "123456"
	validHash, err := hasher.HashPassword(validPassword)
	if err != nil {
		t.Fatalf("error generating hash: %v", err)
	}

	type fields struct {
		config config.IConfig
	}
	type args struct {
		in     string
		inHash string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "CheckPasswordHash success",
			fields: fields{config: configs},
			args: args{
				in:     validPassword,
				inHash: validHash,
			},
			want: true,
		},
		{
			name:   "CheckPasswordHash fail wrong password",
			fields: fields{config: configs},
			args: args{
				in:     "wrong-password",
				inHash: validHash,
			},
			want: false,
		},
		{
			name:   "CheckPasswordHash fail invalid hash",
			fields: fields{config: configs},
			args: args{
				in:     validPassword,
				inHash: "invalid-hash",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{
				config: tt.fields.config,
			}

			if got := h.CheckPasswordHash(tt.args.in, tt.args.inHash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
