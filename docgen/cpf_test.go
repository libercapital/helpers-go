package docgen_test

import (
	"testing"

	"github.com/libercapital/helpers-go/docgen"
)

func TestSeedCPF(t *testing.T) {
	tests := []struct {
		name string
		seed int64
		want string
	}{
		{
			name: "Should generate CPF with seed 2",
			seed: 2,
			want: "76352813488",
		},
		{
			name: "Should generate CPF with seed 123",
			seed: 123,
			want: "50782581757",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := docgen.SeedCPF(tt.seed); got != tt.want {
				t.Errorf("SeedCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomCPF(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Should generate random CPFs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstExec := docgen.RandomCPF()
			secondExec := docgen.RandomCPF()
			if firstExec == secondExec {
				t.Errorf("SeedCPF() = firstExec %v, secondExec %v", firstExec, secondExec)
			}
		})
	}
}
