package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// called automatically
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate a ramdom interger between min and max
func RamdomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //0+min -> max-min+min
}

// generate a ramdom string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RamdomMoney() int64 {
	return RamdomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CNY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
