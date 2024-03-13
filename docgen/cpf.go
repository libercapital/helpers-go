package docgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomCPF() string {
	rand.Seed(time.Now().UnixNano())
	return generateCPF()
}

func SeedCPF(seed int64) string {
	rand.Seed(seed)
	return generateCPF()
}

func generateCPF() string {
	maxDigit := 9

	n1 := rand.Intn(maxDigit)
	n2 := rand.Intn(maxDigit)
	n3 := rand.Intn(maxDigit)
	n4 := rand.Intn(maxDigit)
	n5 := rand.Intn(maxDigit)
	n6 := rand.Intn(maxDigit)
	n7 := rand.Intn(maxDigit)
	n8 := rand.Intn(maxDigit)
	n9 := rand.Intn(maxDigit)

	d1 := n9*2 + n8*3 + n7*4 + n6*5 + n5*6 + n4*7 + n3*8 + n2*9 + n1*10
	d1 = 11 - (d1 % 11)
	if d1 >= 10 {
		d1 = 0
	}

	d2 := d1*2 + n9*3 + n8*4 + n7*5 + n6*6 + n5*7 + n4*8 + n3*9 + n2*10 + n1*11
	d2 = 11 - (d2 % 11)
	if d2 >= 10 {
		d2 = 0
	}

	return strings.ReplaceAll(fmt.Sprint(n1, n2, n3, n4, n5, n6, n7, n8, n9, d1, d2), " ", "")
}
