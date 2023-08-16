package errs

import (
	dgerr "dghire.com/libs/go-common/enums/error"
)

var (
	FakeErr = dgerr.New(500000000, "伪造的错误")

	TokenExpired                     = dgerr.New(4011, "登陆过期，请重新登陆")
	TokenInvalid                     = dgerr.New(4002, "登陆错误，请重新登陆")
	CompanyNotExist                  = dgerr.New(dgerr.SYSTEM_ERROR.Code, "公司不存在")
	CompanyNameExist                 = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "公司名称已存在")
	CompanyDisabled                  = dgerr.New(dgerr.SYSTEM_ERROR.Code, "公司被禁用")
	UserAlreadyInGroup               = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "用户已存在团队内")
	SupplierExist                    = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "供应商已存在")
	GroupNotExist                    = dgerr.New(dgerr.USER_NOT_EXISTS.Code, "团队不存在")
	GroupUserNotExist                = dgerr.New(dgerr.USER_NOT_EXISTS.Code, "非团队内成员")
	UserExist                        = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "用户已存在")
	LinkExpired                      = dgerr.New(dgerr.ARGUMENT_NOT_VALID.Code, "链接已失效，请联系客服重新发送")
	EmptyGroupId                     = dgerr.New(dgerr.ARGUMENT_NOT_VALID.Code, "未选择团队")
	SuperManagerExist                = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "已存在超管账号")
	PlatformUserEditSuperManagerOnly = dgerr.New(dgerr.SYSTEM_ERROR.Code, "平台用户只能操作超管账号")
	AccountOrPasswordErr             = dgerr.New(dgerr.SYSTEM_ERROR.Code, "账号或密码错误")
	SendEmailErr                     = dgerr.New(dgerr.SYSTEM_ERROR.Code, "发送邮件失败")
	EmailAlreadyUsed                 = dgerr.New(dgerr.DUPLICATE_PRIMARY_KEY.Code, "邮箱已被占用")
)
