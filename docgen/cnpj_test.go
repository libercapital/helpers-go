package docgen_test

import (
	"testing"

	"github.com/libercapital/helpers-go/docgen"
)

func TestSeedCNPJ(t *testing.T) {
	tests := []struct {
		name string
		seed int64
		want string
	}{
		{
			name: "Should generate CNPJ with seed 2",
			seed: 2,
			want: "76352813000102",
		},
		{
			name: "Should generate CNPJ with seed 123",
			seed: 123,
			want: "50782581000139",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := docgen.SeedCNPJ(tt.seed); got != tt.want {
				t.Errorf("SeedCNPJ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomCNPJ(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Should generate random CNPJs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstExec := docgen.RandomCNPJ()
			secondExec := docgen.RandomCNPJ()
			if firstExec == secondExec {
				t.Errorf("SeedCNPJ() = firstExec %v, secondExec %v", firstExec, secondExec)
			}
		})
	}
}
