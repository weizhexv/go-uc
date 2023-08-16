package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-uc/pkg/req"
	"go-uc/web/controller"
	"testing"
)

func TestLogin(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewBlankCTX().DgContext
	requestObj := &req.LoginParams{
		Platform: "Web",
		Domain:   "Business",
		Email:    "xuweizhe@reta-inc.com",
		Password: "12345678",
	}
	f := controller.Login.LoginByEmail()
	ret := f(gc, ctx, requestObj)
	fmt.Printf("ret %v", ret)
}
