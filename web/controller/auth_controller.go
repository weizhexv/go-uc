package controller

import (
	dgctx "dghire.com/libs/go-common/context"
	dgerr "dghire.com/libs/go-common/enums/error"
	"dghire.com/libs/go-common/result"
	"dghire.com/libs/go-web/wrapper"
	"github.com/gin-gonic/gin"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/maputil"
	"go-uc/pkg/vo"
)

var Auth = new(auth)

type auth struct {
}

func (a *auth) Bind(rg *gin.RouterGroup) {
	g := rg.Group("/public/auth")

	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[*vo.AuthInfo]]{
		RouterGroup:  g,
		RelativePath: "/verify",
		NonLogin:     true,
		BizHandler:   a.Verify(),
	})
}

func (a *auth) Verify() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[*vo.AuthInfo]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, mp *wrapper.MapRequest) *result.Result[*vo.AuthInfo] {
		ctx := tlog.NewCtx(dc)
		token, ok := maputil.GetString(mp.MP, "token")
		ctx.Infof("auth verfiy token is %s", token)
		if len(token) == 0 || !ok {
			return result.FailByError[*vo.AuthInfo](dgerr.ARGUMENT_NOT_VALID)
		}
		ret, err := service.Auth.Verify(tlog.NewCtx(dc), token)
		ctx.Infof("auth verify ret %v, err %v", ret, err)
		if err != nil {
			return result.FailByError[*vo.AuthInfo](err)
		} else {
			return result.Success(ret)
		}
	}
}
