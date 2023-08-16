package stringutil

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Random(length int) string {
	s := make([]rune, length)
	rand.Seed(time.Now().Unix())
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
