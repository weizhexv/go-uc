package checker

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"go-uc/pkg/req"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/var/roles"
)

func CheckQueryUser(ctx *tlog.Ctx, params *req.QueryUserParams) error {
	if domains.IsNotAccessible(ctx.Domain(), domains.Parse(params.Domain)) {
		return dgerr.NO_PERMISSION
	}
	if ctx.Domain().Is(domains.Business) || ctx.Domain().Is(domains.Supplier) {
		if ctx.DomainId == 0 {
			return dgerr.ARGUMENT_NOT_VALID
		}
		params.DomainId = ctx.DomainId
	}
	//查询B端
	if domains.Equal(params.Domain, domains.Business) {
		return checkQueryBusiness(ctx, params)
	}
	//查询S端
	if domains.Equal(params.Domain, domains.Supplier) {
		return checkQuerySupplier(ctx, params)
	}
	return nil
}

func checkQuerySupplier(ctx *tlog.Ctx, params *req.QueryUserParams) error {
	if ctx.Domain().Is(domains.Platform) {
		if params.Role != roles.SuperManager.Name {
			return errs.PlatformUserEditSuperManagerOnly
		}
	}
	return nil
}

func checkQueryBusiness(ctx *tlog.Ctx, params *req.QueryUserParams) error {
	if ctx.Domain().Is(domains.Platform) {
		if params.Role != roles.SuperManager.Name {
			return errs.PlatformUserEditSuperManagerOnly
		}
		return nil
	} else if ctx.Domain().Is(domains.Business) {
		if ctx.Role == roles.SuperManager.Name {
			//超级管理员能查所有组
			return nil
		} else {
			//非超管只能查看当前组
			if ctx.GroupId == 0 {
				return errs.EmptyGroupId
			}
			gu, err := service.Group.FindGroupUser(ctx, ctx.GroupId, ctx.UserId)
			if err != nil {
				return err
			}
			if gu == nil {
				return errs.GroupUserNotExist
			}
			return nil
		}
	}
	return dgerr.NO_PERMISSION
}

func CheckUpsertUser(ctx *tlog.Ctx, params *req.UpsertUserParams) error {
	if params.Uid == 0 {
		return checkInsertUser(ctx, params)
	} else {
		return checkUpdateUser(ctx, params)
	}
}
func checkInsertUser(ctx *tlog.Ctx, params *req.UpsertUserParams) error {
	//检查参数合法性
	if params == nil {
		return dgerr.ARGUMENT_NOT_VALID
	}
	domain := domains.Parse(params.Domain)
	if len(domain) == 0 {
		ctx.Errorf("params domain is blank")
		return dgerr.ARGUMENT_NOT_VALID
	}
	role := roles.OfName(params.Role)
	if role == nil || role.NotContainsIn(domain) {
		ctx.Errorf("illegal role %s for domain %s", role, domain)
		return dgerr.ARGUMENT_NOT_VALID
	}
	if len(params.Name) == 0 || len(params.Mobile) == 0 || len(params.Email) == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	if ctx.Domain().Is(domains.Business) && ctx.GroupId == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}

	//判断是否有权限
	if domains.IsNotAccessible(ctx.Domain(), domains.Parse(params.Domain)) {
		return dgerr.NO_PERMISSION
	}

	//判断邮箱是否已经注册
	user, err := service.UserImpls.FindByEmail(ctx, params.Email)
	if err != nil {
		return err
	}
	if user != nil {
		return errs.UserExist
	}

	//平台端用户只能编辑超管
	if ctx.Domain().Is(domains.Platform) && domain.IsAny(domains.Business, domains.Supplier) &&
		params.Role != roles.SuperManager.Name {
		return errs.PlatformUserEditSuperManagerOnly
	}

	//如果插入的是超管账号，则判断是否已经有了超管账号
	if params.Role == roles.SuperManager.Name {
		has, err := service.Role.HasSuperManager(ctx, domain, params.DomainId)
		if err != nil {
			return err
		}
		if has {
			return errs.SuperManagerExist
		}
	}
	return nil
}

func checkUpdateUser(ctx *tlog.Ctx, params *req.UpsertUserParams) error {
	//参数校验
	if params == nil {
		return dgerr.ARGUMENT_NOT_VALID
	}
	domain := domains.Parse(params.Domain)
	if len(domain) == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	if params.Uid == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	if len(params.Role) > 0 {
		role := roles.OfName(params.Role)
		if role == nil || role.NotContainsIn(domain) {
			return dgerr.ARGUMENT_NOT_VALID
		}
	}

	//判断是否有操作权限
	if domains.IsNotAccessible(ctx.Domain(), domain) {
		return dgerr.NO_PERMISSION
	}

	//判断该用户是否存在
	user, err := service.UserImpls.FindByUid(ctx, params.Uid)
	if err != nil {
		return err
	}
	if user == nil {
		return dgerr.USER_NOT_EXISTS
	}

	//B端更新用户校验
	if domain.Is(domains.Business) {
		if err = checkUpdateBUser(ctx, params); err != nil {
			return err
		}
	}

	//S端更新用户校验
	if domain.Is(domains.Supplier) {
		if err = checkUpdateSUser(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

func checkUpdateSUser(ctx *tlog.Ctx, params *req.UpsertUserParams) error {
	if ctx.Domain().Is(domains.Platform) {
		if len(params.Role) > 0 && params.Role != roles.SuperManager.Name {
			return errs.PlatformUserEditSuperManagerOnly
		}
	}
	return nil
}

func checkUpdateBUser(ctx *tlog.Ctx, params *req.UpsertUserParams) error {
	//平台用户更新B用户
	if ctx.Domain().Is(domains.Platform) {
		if len(params.Role) > 0 && params.Role != roles.SuperManager.Name {
			return errs.PlatformUserEditSuperManagerOnly
		}
	}
	//B用户更新B用户
	if ctx.Domain().Is(domains.Business) {
		if ctx.Role == roles.SuperManager.Name {
			return nil
		}
		gUser, err := service.Group.FindGroupUser(ctx, ctx.GroupId, ctx.UserId)
		if err != nil {
			return err
		}
		if gUser == nil {
			return dgerr.NO_PERMISSION
		}
		if gUser.RoleId == roles.GroupManager.Id {
			return nil
		}
		if params.Uid != ctx.UserId {
			return dgerr.NO_PERMISSION
		}
	}
	return nil
}
