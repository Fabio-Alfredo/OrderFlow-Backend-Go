package uuid

import (
	"reflect"
	"testing"
)

func TestNewGeneratorUuid(t *testing.T) {
	tests := []struct {
		name string
		want IGenerator
	}{
		{
			name: "NewGeneratorUuid",
			want: &generateUuid{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGeneratorUuid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeneratorUuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateUuid_GenerateId(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "GenerateUuid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &generateUuid{}
			if got := u.GenerateId(); got == "" {
				t.Errorf("GenerateId() = %v, want non-empty string", got)
			}
		})
	}
}
