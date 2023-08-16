package passwords

import (
	"fmt"
	"go-uc/internal/tool/stringutil"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bs), err
}

func Match(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("compare hash and password error: ", err)
		return false
	}
	return true
}

func RandomSeed() string {
	return stringutil.Random(8)
}
