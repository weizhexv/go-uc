package tokens

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/vo"
	"time"
)

type UcClaims struct {
	Uid int64  `json:"uid,omitempty"`
	Dom string `json:"dom,omitempty"`
	Plt string `json:"plt,omitempty"`
	Exp int64  `json:"exp,omitempty"`
}

func (c UcClaims) String() string {
	js, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (c UcClaims) Valid() error {
	if c.Uid <= 0 {
		return errs.TokenInvalid
	}
	if len(c.Dom) == 0 {
		return errs.TokenInvalid
	}
	if len(c.Plt) == 0 {
		return errs.TokenInvalid
	}
	if time.Unix(c.Exp, 0).Before(time.Now()) {
		return errs.TokenExpired
	}
	return nil
}

func Sign(tokenInfo *vo.TokenInfo, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UcClaims{
		Uid: tokenInfo.Uid,
		Dom: tokenInfo.Domain,
		Plt: tokenInfo.Platform,
		Exp: tokenInfo.ExpiresAt.Unix(),
	})

	if signed, err := token.SignedString(secret); err != nil {
		return "", err
	} else {
		return signed, nil
	}
}

func Parse(token string, secret []byte) (*vo.TokenInfo, error) {
	if len(token) == 0 || len(secret) == 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.TokenInvalid
		}
		return secret, nil
	}

	tk, err := jwt.ParseWithClaims(token, &UcClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	c, ok := tk.Claims.(*UcClaims)
	if !ok {
		return nil, errs.TokenInvalid
	}

	return vo.NewTokenInfo(c.Uid, c.Dom, c.Plt, time.Unix(c.Exp, 0)), nil
}

func ParseUnverified(token string) (*UcClaims, error) {
	t, _, err := new(jwt.Parser).ParseUnverified(token, &UcClaims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*UcClaims); ok {
		return claims, nil
	} else {
		return nil, errs.TokenInvalid
	}
}
