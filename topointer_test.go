package bavahelper

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPtrAny(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "string pointer",
			args: args{
				value: "123456",
			},
		},
		{
			name: "int pointer",
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "int64 pointer",
			args: args{
				value: rand.Int63(),
			},
		},
		{
			name: "float pointer",
			args: args{
				value: rand.Float32(),
			},
		},
		{
			name: "float64 pointer",
			args: args{
				value: rand.Float64(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeOf := reflect.TypeOf(PtrAny(tt.args.value))

			if typeOf.Kind() != reflect.Pointer {
				t.Errorf("%s is not a pointer", tt.args.value)
			}
		})
	}
}
