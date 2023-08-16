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
	"go-uc/internal/tool/nonces"
	"go-uc/pkg/req"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/vo"
)

type employee struct {
}

var Employee = new(employee)

func (e *employee) Bind(rg *gin.RouterGroup) {
	g := rg.Group("/public/employee")

	wrapper.Post(&wrapper.RequestHolder[req.SignUpParams, *result.Result[int64]]{
		RouterGroup:  g,
		RelativePath: "/sign-up",
		NonLogin:     true,
		BizHandler:   e.SignUp(),
	})
	wrapper.Post(&wrapper.RequestHolder[req.UpdateEmployeeParams, *result.Result[int64]]{
		RouterGroup:  g,
		RelativePath: "/update",
		BizHandler:   e.Update(),
	})
	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[*vo.EmployeeInfo]]{
		RouterGroup:  g,
		RelativePath: "/info",
		BizHandler:   e.Info(),
	})
	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[map[string]bool]]{
		RouterGroup:  g,
		RelativePath: "/submit-verification-result",
		BizHandler:   e.SubmitVerify(),
	})
	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, string]{
		RouterGroup:  g,
		RelativePath: "/verification-event",
		NonLogin:     true,
		BizHandler:   e.VerificationEvent(),
	})
	wrapper.MapPost(&wrapper.RequestHolder[wrapper.MapRequest, string]{
		RouterGroup:  g,
		RelativePath: "/verification-decision",
		NonLogin:     true,
		BizHandler:   e.VerificationDecision(),
	})
	wrapper.MapGet(&wrapper.RequestHolder[wrapper.MapRequest, *result.Result[bool]]{
		RouterGroup:  g,
		RelativePath: "/exist",
		NonLogin:     true,
		BizHandler:   e.Exist(),
	})
}

func (e *employee) SignUp() wrapper.HandlerFunc[req.SignUpParams, *result.Result[int64]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.SignUpParams) *result.Result[int64] {
		c := tlog.NewCtx(dc)
		c.Infof("employee sign up params: %v", requestObj)
		err := nonces.Check(c, requestObj.Nonce)
		if err != nil {
			return result.FailByError[int64](err)
		}

		info, err := service.UserImpls.FindByEmail(c, requestObj.Email)
		if err != nil {
			return result.FailByError[int64](err)
		}

		if info != nil {
			if info.Domain.Is(domains.Employee) {
				return result.FailByError[int64](errs.UserExist)
			} else {
				return result.FailByError[int64](errs.EmailAlreadyUsed)
			}
		}

		id, err := service.UserImpls.Upsert(c, requestObj.ToInsertUser(c))
		if err != nil {
			c.Errorf("employee sign up error: %v", err)
			return result.FailByError[int64](err)
		}
		c.Infof("employee sign up ret: %d", id)
		err = nonces.Disable(c, requestObj.Nonce)
		if err != nil {
			return result.FailByError[int64](err)
		}
		return result.Success(id)
	}
}

func (e *employee) Update() wrapper.HandlerFunc[req.UpdateEmployeeParams, *result.Result[int64]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *req.UpdateEmployeeParams) *result.Result[int64] {
		ctx := tlog.NewCtx(dc)
		ctx.Infof("update employee params:%v", requestObj)
		if !requestObj.Validate() {
			return result.FailByError[int64](dgerr.ARGUMENT_NOT_VALID)
		}
		if domains.Employee.IsNot(ctx.Domain()) {
			return result.FailByError[int64](dgerr.NO_PERMISSION)
		}
		id, err := service.UserImpls.Upsert(ctx, requestObj.ToUpsertUser(ctx.UserId))
		ctx.Infof("employee update ret %d, err %v", id, err)
		if err != nil {
			return result.FailByError[int64](err)
		} else {
			return result.Success(id)
		}
	}
}

func (e *employee) Info() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[*vo.EmployeeInfo]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[*vo.EmployeeInfo] {
		ctx := tlog.NewCtx(dc)
		if !ctx.Domain().IsAny(domains.Platform, domains.Employee) {
			return result.FailByError[*vo.EmployeeInfo](dgerr.NO_PERMISSION)
		}

		uid := ctx.UserId
		if ctx.Domain().Is(domains.Platform) {
			employeeId, ok := maputil.GetInt64(requestObj.MP, "employeeId")
			if !ok || employeeId == 0 {
				return result.FailByError[*vo.EmployeeInfo](dgerr.ARGUMENT_NOT_VALID)
			}
			uid = employeeId
		}
		ctx.Infof("final request employee uid: %d", uid)

		info, err := service.EmployeeImpl.FindEmployeeInfo(ctx, uid)
		if err != nil {
			return result.FailByError[*vo.EmployeeInfo](err)
		}
		return result.Success(info)
	}
}

func (e *employee) SubmitVerify() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[map[string]bool]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[map[string]bool] {
		ctx := tlog.NewCtx(dc)
		if domains.Employee.IsNot(ctx.Domain()) {
			return result.FailByError[map[string]bool](dgerr.NO_PERMISSION)
		}
		verified, ok := maputil.GetBool(requestObj.MP, "verified")
		if !ok {
			return result.FailByError[map[string]bool](dgerr.ARGUMENT_NOT_VALID)
		}

		sessionId, ok := maputil.GetString(requestObj.MP, "sessionId")
		if !ok {
			return result.FailByError[map[string]bool](dgerr.ARGUMENT_NOT_VALID)
		}

		if err := service.EmployeeImpl.SubmitVerify(ctx, sessionId, verified); err != nil {
			return result.FailByError[map[string]bool](dgerr.SYSTEM_ERROR)
		} else {
			ret := make(map[string]bool)
			ret["verified"] = verified
			return result.Success(ret)
		}
	}
}

func (e *employee) VerificationEvent() wrapper.HandlerFunc[wrapper.MapRequest, string] {
	return func(gc *gin.Context, dc *dgctx.DgContext, _ *wrapper.MapRequest) string {
		signature := gc.GetHeader("X-HMAC-SIGNATURE")
		if signature == "" {
			gc.AbortWithStatus(500)
			return ""
		}

		if cb, ok := gc.Get(gin.BodyBytesKey); ok {
			if cbb, ok := cb.([]byte); ok {
				ctx := tlog.NewCtx(dc)
				err := service.EmployeeImpl.VerificationEvent(ctx, signature, cbb)
				if err != nil {
					gc.AbortWithStatus(500)
				}
			} else {
				gc.AbortWithStatus(500)
			}
		} else {
			gc.AbortWithStatus(500)
		}

		return ""
	}
}

func (e *employee) VerificationDecision() wrapper.HandlerFunc[wrapper.MapRequest, string] {
	return func(gc *gin.Context, dc *dgctx.DgContext, _ *wrapper.MapRequest) string {
		signature := gc.GetHeader("X-HMAC-SIGNATURE")
		if signature == "" {
			gc.AbortWithStatus(500)
			return ""
		}

		if cb, ok := gc.Get(gin.BodyBytesKey); ok {
			if cbb, ok := cb.([]byte); ok {
				ctx := tlog.NewCtx(dc)
				err := service.EmployeeImpl.VerificationDecision(ctx, signature, cbb)
				if err != nil {
					gc.AbortWithStatus(500)
				}
			} else {
				gc.AbortWithStatus(500)
			}
		} else {
			gc.AbortWithStatus(500)
		}

		return ""
	}
}

func (e *employee) Exist() wrapper.HandlerFunc[wrapper.MapRequest, *result.Result[bool]] {
	return func(gc *gin.Context, dc *dgctx.DgContext, requestObj *wrapper.MapRequest) *result.Result[bool] {
		ctx := tlog.NewCtx(dc)
		email, ok := maputil.GetString(requestObj.MP, "email")
		if !ok {
			return result.FailByError[bool](dgerr.ARGUMENT_NOT_VALID)
		}
		exist, err := service.EmployeeImpl.ExistByEmail(ctx, email)
		if err != nil {
			return result.FailByError[bool](err)
		}
		return result.Success(exist)
	}
}
