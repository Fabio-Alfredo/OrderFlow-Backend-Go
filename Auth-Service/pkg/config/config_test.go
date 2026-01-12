package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		viper *viper.Viper
	}
	tests := []struct {
		name string
		args args
		want IConfig
	}{
		{
			name: "Create new config instance",
			args: args{viper: viper.New()},
			want: &config{viper: viper.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.viper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetString(t *testing.T) {

	v := viper.New()
	v.Set("app.name", "TestApp")

	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Get existing string value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "app.name",
			},
			want: "TestApp",
		},
		{
			name: "Get non-existing string value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "non.existing.key",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				viper: tt.fields.viper,
			}
			if got := c.GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetInt(t *testing.T) {
	v := viper.New()
	v.Set("app.port", 8080)

	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Get existing int value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "app.port",
			},
			want: 8080,
		},
		{
			name: "Get non-existing int value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "non.existing.key",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				viper: tt.fields.viper,
			}
			if got := c.GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetBool(t *testing.T) {
	v := viper.New()
	v.Set("feature.enabled", true)

	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Get existing bool value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "feature.enabled",
			},
			want: true,
		},
		{
			name: "Get non-existing bool value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "non.existing.key",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				viper: tt.fields.viper,
			}
			if got := c.GetBool(tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetStringSlice(t *testing.T) {
	v := viper.New()
	v.Set("app.servers", []string{"server1", "server2", "server3"})

	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "Get existing string slice value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "app.servers",
			},
			want: []string{"server1", "server2", "server3"},
		},
		{
			name: "Get non-existing string slice value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "non.existing.key",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				viper: tt.fields.viper,
			}
			if got := c.GetStringSlice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetStringMap(t *testing.T) {
	v := viper.New()
	v.Set("database.host", "localhost")
	v.Set("database.port", 5432)
	v.Set("database.user", "admin")

	type fields struct {
		viper *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "Get existing string map value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "database",
			},
			want: map[string]interface{}{
				"host": "localhost",
				"port": 5432,
				"user": "admin",
			},
		},
		{
			name: "Get non-existing string map value",
			fields: fields{
				viper: v,
			},
			args: args{
				key: "non.existing.key",
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				viper: tt.fields.viper,
			}
			if got := c.GetStringMap(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
