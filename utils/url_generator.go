package utils

import (
	"math/rand"
	"time"
)

func UrlGenerator(length int) string {
	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	newString := make([]rune, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range newString {
		newString[i] = symbols[r.Intn(62)]
	}
	return string(newString)
}
