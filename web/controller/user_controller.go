package controller

import (
	dgctx "dghire.com/libs/go-common/context"
	dgerr "dghire.com/libs/go-common/enums/error"
	"dghire.com/libs/go-common/page"
	"dghire.com/libs/go-common/result"
	"dghire.com/libs/go-web/wrapper"
	"github.com/gin-gonic/gin"
	"go-uc/internal/checker"
	"go-uc/internal/model"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/maputil"
	"go-uc/internal/tool/nonces"
	"go-uc/internal/tool/passwords"
	"go-uc/pkg/req"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
	"go/types"
	"math"
	"strconv"
	"strings"
)

type user struct {
}

var User = new(user)

func (u *user) Bind(rg *gin.RouterGroup) {
	g := rg.Group("/public/user")

	wrapper.Get(&wrapper.RequestHolder[req.QueryUserParams, *result.Result[*page.PageList[vo.UserDetail]]]{
		RouterGroup:  g,
		RelativePath: "/query",
		BizHandler:   u.Query(),
	})
	wrapper.Post(&wrapper.RequestHolder[req.UpsertUserParams, *result.Result[int64]]{
		RouterGroup:  g,
		RelativePath: "/upsert",
		BizHandler:   u.Upsert(),
	})
	wrapper.Post(&wrapper.RequestHolder[req.SetPasswordParams, *result.Result[types.Nil]]{
		RouterGroup:  g,
		RelativePath: "/set/password",
		NonLogin:     true,
		BizHandler:   u.SetPassword(),
	})
	wrapper.Post(&wrapper.RequestHolder[req.ResetPasswordParams, *result.Result[types.Nil]]{
		RouterGroup:  g,
		RelativePath: "/reset/password",
		BizHandler:   u.ResetPassword(),
	})
	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[types.Nil]]{
		RouterGroup:  g,
		RelativePath: "/reset/password/email",
		BizHandler:   u.ResetPasswordByEmail(),
	})
	wrapper.Get(&wrapper.RequestHolder[types.Nil, *result.Result[*vo.UserInfo]]{
		RouterGroup:  g,
		RelativePath: "/info/:uid",
		BizHandler:   u.Info(),
	})
	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[*vo.UserDetail]]{
		RouterGroup:  g,
		RelativePath: "/info",
		BizHandler:   u.InfoV2(),
	})
	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[[]*vo.IDNamePair]]{
		RouterGroup:  g,
		RelativePath: "/filter",
		BizHandler:   u.Filter(),
	})
	wrapper.Get(&wrapper.RequestHolder[req.UserFilterParams, *result.Result[[]*vo.IDNameEmailTriple]]{
		RouterGroup:  g,
		RelativePath: "/filter/v2",
		BizHandler:   u.FilterV2(),
	})
}

func (u *user) Query() wrapper.HandlerFunc[req.QueryUserParams, *result.Result[*page.PageList[vo.UserDetail]]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.QueryUserParams) *result.Result[*page.PageList[vo.UserDetail]] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("user query params %v", requestObj)
		err := checker.CheckQueryUser(ctx, requestObj)
		if err != nil {
			return result.FailByError[*page.PageList[vo.UserDetail]](err)
		}
		ret, err := service.UserImpls.Page(ctx, requestObj.ToUserQuery(ctx))
		ctx.Infof("user query ret %v, err %v", ret, err)
		if err != nil {
			return result.FailByError[*page.PageList[vo.UserDetail]](err)
		}
		return result.Success(ret)
	}
}

func (u *user) Upsert() wrapper.HandlerFunc[req.UpsertUserParams, *result.Result[int64]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.UpsertUserParams) *result.Result[int64] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("user upsert params %v", requestObj)
		err := checker.CheckUpsertUser(ctx, requestObj)
		if err != nil {
			return result.FailByError[int64](err)
		}
		upsertUser, err := requestObj.ToUpsertUser(ctx)
		if err != nil {
			return result.FailByError[int64](err)
		}
		uid, err := service.UserImpls.Upsert(ctx, upsertUser)
		ctx.Infof("user upsert ret %v, err %v", uid, err)
		if err != nil {
			return result.FailByError[int64](err)
		}
		return result.Success(uid)
	}
}

func (u *user) SetPassword() wrapper.HandlerFunc[req.SetPasswordParams, *result.Result[types.Nil]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.SetPasswordParams) *result.Result[types.Nil] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("set password params %v", requestObj)
		info, err := service.UserImpls.FindByEmail(ctx, requestObj.Email)
		if err != nil || info == nil {
			return result.FailByError[types.Nil](dgerr.USER_NOT_EXISTS)
		}
		err = nonces.Check(ctx, requestObj.Nonce)
		if err != nil {
			return result.FailByError[types.Nil](err)
		}
		err = service.UserImpls.ResetPassword(ctx, requestObj.Password, info.Uid)
		if err != nil {
			ctx.Infof("user set password err: %v", err)
			return result.FailByError[types.Nil](err)
		}
		if err = nonces.Disable(ctx, requestObj.Nonce); err != nil {
			return result.FailByError[types.Nil](err)
		} else {
			return result.Success(types.Nil{})
		}
	}
}

func (u *user) ResetPassword() wrapper.HandlerFunc[req.ResetPasswordParams, *result.Result[types.Nil]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.ResetPasswordParams) *result.Result[types.Nil] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("reset password params %v", requestObj)
		info, err := service.UserImpls.FindByEmail(ctx, requestObj.Email)
		if err != nil || info == nil {
			return result.FailByError[types.Nil](dgerr.USER_NOT_EXISTS)
		}
		if ctx.UserId != info.Uid {
			return result.FailByError[types.Nil](dgerr.NO_PERMISSION)
		}
		hash, err := service.UserImpls.FindPwdByUid(ctx, info.Uid)
		if err != nil {
			return result.FailByError[types.Nil](errs.AccountOrPasswordErr)
		}
		if len(requestObj.OldPassword) > 0 && !passwords.Match(requestObj.OldPassword, hash) {
			return result.FailByError[types.Nil](errs.AccountOrPasswordErr)
		}
		err = service.UserImpls.ResetPassword(ctx, requestObj.Password, info.Uid)
		if err != nil {
			ctx.Infof("user reset password err: %v", err)
			return result.FailByError[types.Nil](err)
		}
		return result.Success(types.Nil{})
	}
}

func (u *user) ResetPasswordByEmail() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[types.Nil]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[types.Nil] {
		ctx := tlog.NewCtx(dc)
		email, ok := maputil.GetString(requestObj.MP, "email")
		ctx.Infof("reset password by email: %s", email)
		if !ok || len(email) == 0 {
			return result.FailByError[types.Nil](dgerr.ARGUMENT_NOT_VALID)
		}
		info, err := service.UserImpls.FindByEmail(ctx, email)
		if err != nil || info == nil {
			return result.FailByError[types.Nil](dgerr.USER_NOT_EXISTS)
		}
		if domains.IsNotAccessible(ctx.Domain(), info.Domain) ||
			!roles.IsAny(roles.OfName(ctx.Role), roles.SuperManager, roles.Manger, roles.GroupManager) {
			return result.FailByError[types.Nil](dgerr.NO_PERMISSION)
		}
		err = service.Email.SendResetPasswordEmail(ctx, info.Uid)
		if err != nil {
			ctx.Infof("user reset password by email err: %v", err)
			return result.FailByError[types.Nil](errs.SendEmailErr)
		}
		return result.Success(types.Nil{})
	}
}

func (u *user) Info() wrapper.HandlerFunc[types.Nil, *result.Result[*vo.UserInfo]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *types.Nil) *result.Result[*vo.UserInfo] {
		ctx := tlog.NewCtx(dc)
		uid, err := strconv.ParseInt(gc.Param("uid"), 10, 64)
		if err != nil || uid <= 0 {
			return result.FailByError[*vo.UserInfo](dgerr.ARGUMENT_NOT_VALID)
		}
		info, err := u.getUserInfo(ctx, uid)
		ctx.Infof("user info ret %v, err %v", info, err)
		if err != nil {
			return result.FailByError[*vo.UserInfo](err)
		}
		return result.Success(info)
	}
}

func (u *user) InfoV2() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[*vo.UserDetail]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[*vo.UserDetail] {
		ctx := tlog.NewCtx(dc)
		uid, ok := maputil.GetInt64(requestObj.MP, "uid")
		if !ok {
			return result.FailByError[*vo.UserDetail](dgerr.ARGUMENT_NOT_VALID)
		}
		detail, err := u.getUserDetail(ctx, uid)
		ctx.Infof("user info v2 ret %v, err %v", detail, err)
		if err != nil {
			return result.FailByError[*vo.UserDetail](err)
		}
		return result.Success(detail)
	}
}

func (u *user) getUserInfo(ctx *tlog.Ctx, uid int64) (*vo.UserInfo, error) {
	if uid <= 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}

	info, err := service.UserImpls.FindByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (u *user) getUserDetail(ctx *tlog.Ctx, uid int64) (*vo.UserDetail, error) {
	info, err := u.getUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	dmInfo, err := service.DomainService.GetDomainInfo(ctx, info.Domain, info.DomainId)
	if err != nil {
		return nil, err
	}
	return vo.UserDetailOf(nil, info, dmInfo.DomainName), nil
}

func (u *user) Filter() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[[]*vo.IDNamePair]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[[]*vo.IDNamePair] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("user filter params %v", requestObj.MP)
		//解析参数
		domain, ok := maputil.GetString(requestObj.MP, "domain")
		if !ok && len(domain) == 0 {
			return result.FailByError[[]*vo.IDNamePair](dgerr.ARGUMENT_NOT_VALID)
		}
		//如果用户是B用户或者S用户只能查看本公司信息
		domainId, _ := maputil.GetInt64(requestObj.MP, "domainId")
		if ctx.Domain().IsAny(domains.Business, domains.Supplier) {
			domainId = ctx.DomainId
		}
		//判断权限
		if domains.IsNotAccessible(ctx.Domain(), domains.Parse(domain)) {
			return result.FailByError[[]*vo.IDNamePair](dgerr.NO_PERMISSION)
		}

		infos, err := service.UserImpls.FindByDomainId(ctx, domains.Parse(domain), domainId)
		ctx.Infof("user filter ret %v, err %v", infos, err)
		if err != nil {
			return result.FailByError[[]*vo.IDNamePair](err)
		}
		return result.Success(infosToPairs(infos))
	}
}

func infosToPairs(infos []*vo.UserInfo) []*vo.IDNamePair {
	if len(infos) == 0 {
		return []*vo.IDNamePair{}
	}
	pairs := make([]*vo.IDNamePair, len(infos))
	if len(infos) == 0 {
		return pairs
	}
	i := 0
	for _, info := range infos {
		pairs[i] = vo.IDNamePairOf(info.Uid, info.Name)
		i++
	}
	return pairs
}

func (u *user) FilterV2() wrapper.HandlerFunc[req.UserFilterParams, *result.Result[[]*vo.IDNameEmailTriple]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.UserFilterParams) *result.Result[[]*vo.IDNameEmailTriple] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("user filter v2 params %v", requestObj)
		domain := domains.Parse(requestObj.Domain)

		if domains.IsNotAccessible(ctx.Domain(), domain) {
			return result.FailByError[[]*vo.IDNameEmailTriple](dgerr.NO_PERMISSION)
		}
		//如果用户是B用户或者S用户只能查看本公司信息
		if ctx.Domain().IsAny(domains.Business, domains.Supplier) {
			requestObj.DomainId = ctx.DomainId
		}

		infos, err := service.UserImpls.Query(ctx, &model.UserQuery{
			Domain:       domain,
			DomainId:     requestObj.DomainId,
			CheckEnabled: true,
			Enabled:      true,
			Roles:        u.parseRoles(requestObj.Role),
			GroupRole:    roles.OfName(requestObj.GroupRole),
			GroupId:      requestObj.GroupId,
			PageNo:       1,
			PageSize:     math.MaxInt,
		})
		ctx.Infof("user filter v2 ret %v, err %v", infos, err)
		if err != nil {
			return result.FailByError[[]*vo.IDNameEmailTriple](err)
		}

		return result.Success(infosToTriples(infos))
	}
}

func (u *user) parseRoles(roleString string) []*roles.Role {
	var rs []*roles.Role
	if len(roleString) > 0 {
		arr := strings.Split(roleString, ",")
		for _, s := range arr {
			if r := roles.OfName(s); r != nil {
				rs = append(rs, r)
			}
		}
	}
	return rs
}

func infosToTriples(infos []*vo.UserInfo) []*vo.IDNameEmailTriple {
	if len(infos) == 0 {
		return []*vo.IDNameEmailTriple{}
	}
	triples := make([]*vo.IDNameEmailTriple, len(infos))
	if len(infos) == 0 {
		return triples
	}
	i := 0
	for _, info := range infos {
		triples[i] = vo.IDNameEmailTripleOf(info.Uid, info.Name, info.Email)
		i++
	}
	return triples
}
