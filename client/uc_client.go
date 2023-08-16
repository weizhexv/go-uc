package client

import (
	dgctx "dghire.com/libs/go-common/context"
	dgerr "dghire.com/libs/go-common/enums/error"
	"dghire.com/libs/go-common/result"
	"dghire.com/libs/go-web/wrapper"
	"github.com/gin-gonic/gin"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/maputil"
	"go-uc/internal/tool/typeutil"
	"go-uc/pkg/vo"
)

type ucClient struct {
}

var UcClient = new(ucClient)

func (uc *ucClient) Bind(rg *gin.RouterGroup) {
	g := rg.Group("/internal/user")

	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[[]*vo.UserInfo]]{
		RouterGroup:  g,
		NonLogin:     true,
		RelativePath: "/infos",
		BizHandler:   uc.GetUserInfos,
	})

	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[*vo.UserInfo]]{
		RouterGroup:  g,
		NonLogin:     true,
		RelativePath: "/info",
		BizHandler:   uc.GetUserinfo,
	})
}

func (uc *ucClient) GetUserInfos(gc *gin.Context, dc *dgctx.DgContext, obj *wrapper.MapRequest) *result.Result[[]*vo.UserInfo] {
	ctx := tlog.NewCtx(dc)
	uids, err := typeutil.ParseInt64s(obj.MP["uids"])
	if err != nil {
		ctx.Errorf("parse int64 uids error: %v", err)
		return result.FailByError[[]*vo.UserInfo](dgerr.ARGUMENT_NOT_VALID)
	}
	if len(uids) == 0 {
		return result.Success([]*vo.UserInfo{})
	}

	infos, err := service.UserImpls.FindByUids(ctx, uids)
	if err != nil {
		ctx.Errorf("find user by uids error: %v", err)
		return result.FailByError[[]*vo.UserInfo](dgerr.ARGUMENT_NOT_VALID)
	}
	return result.Success(infos)
}

func (uc *ucClient) GetUserinfo(gc *gin.Context, dc *dgctx.DgContext, obj *wrapper.MapRequest) *result.Result[*vo.UserInfo] {
	c := tlog.NewCtx(dc)
	uid, ok := maputil.GetInt64(obj.MP, "uid")
	if !ok {
		return result.FailByError[*vo.UserInfo](dgerr.ARGUMENT_NOT_VALID)
	}

	info, err := service.UserImpls.FindByUid(c, uid)
	if err != nil {
		c.Errorf("get user info err: %v", err)
		return result.FailByError[*vo.UserInfo](err)
	}

	return result.Success(info)
}
