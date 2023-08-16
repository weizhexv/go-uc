package controller

import (
	dgctx "dghire.com/libs/go-common/context"
	"dghire.com/libs/go-common/result"
	"dghire.com/libs/go-web/wrapper"
	"github.com/gin-gonic/gin"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"go-uc/pkg/req"
	"go-uc/pkg/var/headers"
	"go-uc/pkg/vo"
	"go-uc/vconfig"
	"go/types"
)

var Login = new(login)

type login struct {
}

func (l *login) Bind(rg *gin.RouterGroup) {
	g := rg.Group("/public")

	wrapper.Post(&wrapper.RequestHolder[req.LoginParams, *result.Result[*vo.LoginDetail]]{
		RouterGroup:  g,
		RelativePath: "/login/email",
		NonLogin:     true,
		BizHandler:   l.LoginByEmail(),
	})

	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[types.Nil]]{
		RouterGroup:  g,
		RelativePath: "/logout",
		BizHandler:   l.Logout(),
	})
}

func (l *login) LoginByEmail() wrapper.HandlerFunc[req.LoginParams, *result.Result[*vo.LoginDetail]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.LoginParams) *result.Result[*vo.LoginDetail] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("login by email params %v", requestObj)
		loginCtx, err := requestObj.ToLoginContext()
		if err != nil {
			return result.FailByError[*vo.LoginDetail](err)
		}
		service.Login.Do(ctx, loginCtx)
		if !loginCtx.Success {
			return result.FailByError[*vo.LoginDetail](loginCtx.Err)
		}
		detail, err := loginCtx.ToLoginDetail()
		if err != nil {
			ctx.Errorf("login by email err: %v", err)
			return result.FailByError[*vo.LoginDetail](err)
		}
		ctx.Infof("login by email ret %v", detail)
		l.setCookies(gc, detail.AuthToken)
		return result.Success(detail)
	}
}

func (l *login) Logout() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[types.Nil]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[types.Nil] {
		l.deleteCookies(gc)
		return result.Success(types.Nil{})
	}
}

func (l *login) setCookies(ctx *gin.Context, token string) {
	ctx.SetCookie(headers.Token, token, 60*60*24*20, "/", vconfig.HostCookie(), false, false)
}

func (l *login) deleteCookies(ctx *gin.Context) {
	ctx.SetCookie(headers.Token, "", 0, "/", vconfig.HostCookie(), false, false)
}
