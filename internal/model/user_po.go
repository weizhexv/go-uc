package model

import (
	"go-uc/pkg/var/domains"
)

type UserModel struct {
	domains.Domain
	Id       int64
	Uid      int64
	Password string
	Seed     []byte
}
