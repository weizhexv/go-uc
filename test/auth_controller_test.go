package test

import (
	"dghire.com/libs/go-web/wrapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-uc/web/controller"
	"testing"
)

func TestVerify(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewBlankCTX().DgContext
	mp := &wrapper.MapRequest{MP: map[string]any{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI0LCJkbW4iOiJCVVNJTkVTUyIsInBsdCI6IldFQiIsImV4cCI6MTY3OTcwNzU1M30.vBNqsw--nw3D1DAiAjlSP62r-zHzROxmAN_pO34L31Q"}}
	f := controller.Auth.Verify()
	ret := f(gc, ctx, mp)
	fmt.Printf("ret %v", ret)
	assert.True(t, ret.Success)
}
