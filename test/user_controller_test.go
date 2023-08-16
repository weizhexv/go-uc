package test

import (
	"dghire.com/libs/go-common/page"
	"dghire.com/libs/go-web/wrapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-uc/pkg/req"
	"go-uc/web/controller"
	"testing"
)

func TestUpsertBusinessUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Domain:   "BUSINESS",
		DomainId: 1,
		Name:     "伟小牛",
		Mobile:   "13521511968",
		Email:    "business@reta-inc.com",
		Role:     "GROUP_MANAGER",
		Enabled:  true,
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestInsertBusinessUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Domain:   "BUSINESS",
		DomainId: 1,
		Name:     "伟商业哲学",
		Mobile:   "13521511968",
		Email:    "business6@reta-inc.com",
		Role:     "FINANCE",
		Enabled:  true,
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestPlatformInsertBusinessUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewPlatformUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Domain:   "BUSINESS",
		DomainId: 1,
		Name:     "伟黯然伤神",
		Mobile:   "13521511968",
		Role:     "SUPER_MANAGER",
		Enabled:  true,
		Uid:      32,
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestUpsertPlatformUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewPlatformUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Domain:  "PLATFORM",
		Name:    "伟平台",
		Mobile:  "13521511968",
		Email:   "admin@reta-inc.com",
		Role:    "MANAGER",
		Enabled: true,
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestInsertSupplierUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewPlatformUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Domain:   "SUPPLIER",
		DomainId: 1,
		Name:     "伟供应",
		Mobile:   "13521511968",
		Email:    "weizhexv@163.com",
		Role:     "SUPER_MANAGER",
		Enabled:  true,
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestUpdateBusinessUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewPlatformUserCTX().DgContext
	requestObj := &req.UpsertUserParams{
		Uid:    24,
		Domain: "BUSINESS",
		Name:   "伟小牛惦记你",
		Mobile: "13521511988",
		Role:   "SUPER_MANAGER",
	}
	f := controller.User.Upsert()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %+v", ret)
	assert.True(t, ret.Success)
}

func TestQueryUser(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &req.QueryUserParams{
		Domain:   "BUSINESS",
		DomainId: 1,
		Enabled:  "true",
		PageParam: page.PageParam{
			PageNo:   1,
			PageSize: 10,
		},
	}

	f := controller.User.Query()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestResetPassword(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &req.ResetPasswordParams{
		Email:       "xuweizhe@reta-inc.com",
		Password:    "12345678",
		OldPassword: "12345678",
	}
	f := controller.User.ResetPassword()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestResetPasswordByEmail(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &wrapper.MapRequest{
		MP: make(map[string]any),
	}
	requestObj.MP["email"] = "xuweizhe@reta-inc.com"
	f := controller.User.ResetPasswordByEmail()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestSetPasswordByNonce(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBlankCTX().DgContext
	requestObj := &req.SetPasswordParams{
		Email:    "xuweizhe@reta-inc.com",
		Password: "12345678",
		Nonce:    "48_KXBL6uASzEKxRgfYfPKI",
	}
	f := controller.User.SetPassword()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestUserInfo(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &wrapper.MapRequest{
		MP: make(map[string]any),
	}
	requestObj.MP["uid"] = 35
	f := controller.User.InfoV2()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestUserFilter(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &wrapper.MapRequest{
		MP: make(map[string]any),
	}
	requestObj.MP["domain"] = "BUSINESS"
	f := controller.User.Filter()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestPlatformUserFilter(t *testing.T) {
	gc := new(gin.Context)
	dc := NewPlatformUserCTX().DgContext
	requestObj := &wrapper.MapRequest{
		MP: make(map[string]any),
	}
	requestObj.MP["domain"] = "BUSINESS"
	requestObj.MP["domainId"] = 1

	f := controller.User.Filter()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestUserFilterV2(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBusinessUserCTX().DgContext
	requestObj := &req.UserFilterParams{
		Domain:    "BUSINESS",
		Role:      "SUPER_MANAGER",
		GroupRole: "HR",
		GroupId:   1,
	}
	f := controller.User.FilterV2()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}

func TestBatchGetUserInfo(t *testing.T) {
	gc := new(gin.Context)
	dc := NewBlankCTX().DgContext
	requestObj := &wrapper.MapRequest{
		MP: make(map[string]any),
	}
	requestObj.MP["uids"] = "1,2,38"

	f := controller.Internal.Infos()
	ret := f(gc, dc, requestObj)
	fmt.Printf("ret %v", ret)
}
