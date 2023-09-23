package utils

import "math/rand"

var runes = []rune(
	"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
)

func RandomURL(size int) string {
	str := make([]rune, size)

	for i := range str {
		// rand.Intn(len(runes))  --> this selects a random index or number from length of runes array
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}
