package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/tokens"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/vo"
)

var Auth = new(auth)

type auth struct {
}

func (a *auth) Verify(c *tlog.Ctx, token string) (*vo.AuthInfo, error) {
	c.Infof("auth verify by token: %s", token)

	claims, err := tokens.ParseUnverified(token)
	if err != nil {
		c.Error("parse unverified err:", err)
		return nil, err
	}
	c.Infof("token claims: %v", claims)

	usr, err := UserImpls.FindByUid(c, claims.Uid)
	c.Infof("auth verfiy find user result: %v", usr)

	if err != nil {
		c.Error("find user err:", err)
		return nil, dgerr.USER_NOT_EXISTS
	}

	if !usr.Enabled || usr.Deleted {
		c.Error("disabled user, uid:", claims.Uid)
		return nil, dgerr.DISABLED_USER
	}

	if usr.Domain == domains.Business {
		cus, err := Customer.FindById(c, usr.DomainId)
		if err != nil {
			c.Error("find customer err:", err)
			return nil, errs.CompanyNotExist
		}
		if cus == nil {
			c.Error("customer not exist")
			return nil, errs.CompanyNotExist
		}
	}

	seed, err := UserImpls.FindSeedByUid(c, claims.Uid)
	if err != nil {
		c.Error("find user seed err:", err)
		return nil, err
	}

	tokenInfo, err := tokens.Parse(token, seed)
	if err != nil {
		c.Error("parse token err:", err)
		return nil, err
	}

	return vo.NewAuthInfo(tokenInfo, usr), nil
}
