package util

import(
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int) int {
	return min + rand.Intn(max - min + 1)
}

func RandomString(n int) string {
	b := make([]byte, n)
	k:= len(alphabet)

	for i := range b {
		b[i] = alphabet[rand.Intn(k)]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}