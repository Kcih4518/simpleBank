package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// Seed the random number generator with the current time.
	// If not seeded, the generator will produce the same sequence of numbers
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random number in the range [min, max]
func RandomInt(min, max int64) int64 {
	// rand.Int63n(max-min+1) returns a random number in the range [0, max-min+1)
	// min + rand.Int63n(max-min+1) returns a random number in the range [min, max]
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
