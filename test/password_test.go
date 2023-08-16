package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-uc/internal/tool/passwords"
	"testing"
)

func TestPasswordTool(t *testing.T) {
	password := "secret"
	hash, err := passwords.Hash(password)
	assert.True(t, err == nil)

	hash2, err := passwords.Hash(password)
	assert.True(t, err == nil)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	fmt.Println("Hash2:   ", hash2)

	match := passwords.Match(password, hash)
	fmt.Println("Match:   ", match)
	assert.True(t, match)
	match = passwords.Match(password, hash2)
	fmt.Println("Match:   ", match)
	assert.True(t, match)
	match = passwords.Match(password, "$2a$10$ao8Jt9F6vVBvyqns2qsnNuNlB5TXDabrI7ucFGITA9xbgw96i.GxJ")
	fmt.Println("Match:   ", match)
	assert.False(t, match)
	match = passwords.Match("secrett", hash2)
	fmt.Println("Match:   ", match)
	assert.False(t, match)
}

func TestDecodeJavaHash(t *testing.T) {
	ok := passwords.Match("secretsssss", "$2a$10$My6.6IZ3su95SJHgNykKdOXQ.TnSBkvAFZhHRTT/Ffa65ujgRwtV6")
	assert.True(t, ok)
}

func TestSeed(t *testing.T) {
	sd := passwords.RandomSeed()
	fmt.Println(sd)
}
