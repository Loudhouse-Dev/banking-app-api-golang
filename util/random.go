package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Generate a random number between min and max
func RandomInt(min, max int64) int64 {
	// rand.Int63n returns a random number (int64) between min and max
	return min + rand.Int63n(max - min + 1) 
}

// Generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generate a random account owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generate a random account balance
func RandomMoney() int64 {
	return RandomInt(0, 5000)
}

// Generate a random account currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}