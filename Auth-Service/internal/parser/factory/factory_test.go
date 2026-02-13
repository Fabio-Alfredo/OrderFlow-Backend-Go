package factory

import (
	"Auth-Service/internal/parser"
	"Auth-Service/pkg/config"
	"reflect"
	"testing"
)

func TestNewParserFactory(t *testing.T) {
	tests := []struct {
		name string
		want parser.IFactory
	}{
		{
			name: "Test NewParserFactory",
			want: &parserFactory{
				parsers: map[string]parser.IParser{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParserFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserFactory_Set(t *testing.T) {
	configs, _ := config.Load("../../../../")
	newParser := parser.NewUserDomainToUserRepositoryParser(configs)
	type fields struct {
		parsers map[string]parser.IParser
	}
	type args struct {
		key    string
		parser parser.IParser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test set parser success",
			fields: fields{
				parsers: map[string]parser.IParser{},
			},
			args: args{
				key:    parser.UserDomainToUserRepositoryParser,
				parser: newParser,
			},
			wantErr: false,
		},
		{
			name: "Test set parser error in key empty",
			fields: fields{
				parsers: map[string]parser.IParser{},
			},
			args: args{
				key: "",
			},
			wantErr: true,
		},
		{
			name: "Test set parser error in parser empty",
			fields: fields{
				parsers: map[string]parser.IParser{},
			},
			args: args{
				key:    parser.UserDomainToUserRepositoryParser,
				parser: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &parserFactory{
				parsers: tt.fields.parsers,
			}
			if err := f.Set(tt.args.key, tt.args.parser); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
