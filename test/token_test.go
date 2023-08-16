package test

import (
	"errors"
	"fmt"
	"go-uc/internal/tool/tokens"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/vo"
	"testing"
	"time"
)

var secret = []byte("3FXcDeqk")

func TestSignAndParse(t *testing.T) {
	tokenInfo := vo.NewTokenInfo(10, string(domains.Business), "WEB", time.Now().Add(time.Minute*5))

	token, err := tokens.Sign(tokenInfo, secret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)

	parse, err := tokens.Parse(token, secret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("parse: %v\n", parse)

	if parse.Uid != tokenInfo.Uid {
		t.Fatal("uid mismatch")
	}
	if parse.Domain != tokenInfo.Domain {
		t.Fatal("tokenInfo mismatch")
	}
	if parse.Platform != tokenInfo.Platform {
		t.Fatal("platform mismatch")
	}
	if parse.ExpiresAt != tokenInfo.ExpiresAt {
		t.Fatal("expiresAt mismatch")
	}
}

func TestSignAndParseExpired(t *testing.T) {
	tokenInfo := vo.NewTokenInfo(10, string(domains.Business), "WEB", time.Now().Add(-time.Second*5))

	token, err := tokens.Sign(tokenInfo, secret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)

	_, err = tokens.Parse(token, secret)
	tokenExpiredErr := errs.TokenExpired
	if err != nil {
		if errors.Is(err, tokenExpiredErr) {
			fmt.Println("receive token expired error")
			return
		}
	}
	t.Fatal("expect token expired error")
}

func TestParseUnverified(t *testing.T) {
	tokenInfo := vo.NewTokenInfo(10, string(domains.Business), "WEB", time.Now().Add(time.Minute*5))

	token, err := tokens.Sign(tokenInfo, secret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %s\n", token)

	claims, err := tokens.ParseUnverified(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("claims: %v\n", claims)
}
