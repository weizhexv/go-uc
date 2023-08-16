package test

import (
	"github.com/stretchr/testify/assert"
	"go-uc/internal/tool/emails"
	"go-uc/pkg/vo"
	"testing"
	"time"
)

func TestSendEmail(t *testing.T) {
	ctx := NewPlatformUserCTX()
	detail := &vo.UserDetail{
		UserInfo: &vo.UserInfo{
			Uid:      1,
			Name:     "xuweizhe",
			Email:    "996997256@qq.com",
			Domain:   "BUSINESS",
			DomainId: 1,
		},
		DomainName: "D & G",
		Me:         false,
	}
	err := emails.SendRegisterEmail(ctx, detail)
	assert.True(t, err == nil)

	time.Sleep(time.Second * 10)
}
