package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"errors"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/passwords"
	"go-uc/pkg/var/platforms"
	"go-uc/pkg/vo"
	"time"
)

const day = time.Hour * time.Duration(24)

type login struct {
}

var Login = new(login)

func (l *login) Do(c *tlog.Ctx, lc *model.LoginContext) {
	//check user exist or not
	info, err := l.checkUser(c, lc)
	if err != nil {
		c.Errorf("login check user error: %v", err)
		lc.Abort(err)
		return
	}

	//check password
	hash, err := UserImpls.FindPwdByUid(c, info.Uid)
	if err != nil {
		c.Errorf("find pwd by uid error: %v", err)
		lc.Abort(dgerr.WRONG_PASSWORD)
		return
	}

	if !passwords.Match(lc.Pwd, hash) {
		c.Warnf("login passowrd not match, pwd: %s, hash: %s", lc.Pwd, hash)
		lc.Abort(dgerr.WRONG_PASSWORD)
		return
	}

	//refresh login
	uLogin, err := l.do(c, lc, info)
	if err != nil {
		c.Errorf("do login error: %v", err)
		lc.Abort(dgerr.LOGIN_ERROR)
		return
	}

	// set domain name and status
	dmInfo, err := DomainService.GetDomainInfo(c, info.Domain, info.DomainId)
	if err != nil {
		c.Errorf("get domain info error: %v", err)
		lc.Abort(err)
		return
	}

	//set user group info
	gInfos, err := l.findGroupUserInfos(c, info.Uid)
	if err != nil {
		c.Errorf("find group user infos error: %v", err)
		lc.Abort(err)
		return
	}

	//get seed
	seed, err := UserImpls.FindSeedByUid(c, info.Uid)
	if err != nil {
		c.Errorf("find seed by uid error: %v", err)
		lc.Abort(err)
		return
	}

	lc.Done(info, uLogin.ToInfo(), dmInfo, gInfos, seed)
}

func (l *login) checkUser(c *tlog.Ctx, lc *model.LoginContext) (*vo.UserInfo, error) {
	usr, err := UserImpls.FindByEmail(c, lc.Email)
	if err != nil {
		return nil, err
	}
	if usr == nil || usr.Domain != lc.Domain {
		return nil, dgerr.USER_NOT_EXISTS
	}
	return usr, nil
}

func (l *login) do(c *tlog.Ctx, lc *model.LoginContext, info *vo.UserInfo) (*dal.UserLogin, error) {
	return daog.AutoTransWithResult[*dal.UserLogin](func() (*daog.TransContext, error) {
		return daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	}, func(tc *daog.TransContext) (*dal.UserLogin, error) {
		uLogin, err := l.findLogin(tc, lc)
		if err != nil {
			c.Errorf("find login error: %v", err)
			return nil, err
		}

		var ret *dal.UserLogin
		if uLogin == nil {
			ret, err = l.createLogin(tc, info, lc.Platform)
			if err != nil {
				c.Errorf("create login err: %v", err)
				return nil, err
			}
		} else {
			ret, err = l.refreshLogin(tc, uLogin)
			if err != nil {
				c.Errorf("refresh login err: %v", err)
				return nil, err
			}
		}
		return ret, err
	})
}

func (l *login) findLogin(tc *daog.TransContext, lc *model.LoginContext) (*dal.UserLogin, error) {
	mc := daog.NewMatcher().Eq(dal.UserLoginFields.Domain, lc.Domain.String())
	mc.Eq(dal.UserLoginFields.Platform, lc.Platform.String())
	mc.Eq(dal.UserLoginFields.Identity, lc.Email)
	return repo.Login.QueryOneMatcher(tc, mc)
}

func (l *login) createLogin(tc *daog.TransContext, user *vo.UserInfo, platform platforms.Platform) (*dal.UserLogin, error) {
	ulogin := dal.NewUserLogin(user, platform, day*30)
	rows, err := repo.Login.Insert(tc, ulogin)
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, errors.New("create login affect rows != 1")
	}
	return ulogin, nil
}

func (l *login) refreshLogin(tc *daog.TransContext, uLogin *dal.UserLogin) (*dal.UserLogin, error) {
	uLogin.ExpiredAt = ttypes.NormalDatetime(time.Now().Add(day * 30))
	uLogin.ModifiedAt = ttypes.NormalDatetime(time.Now())

	rows, err := repo.Login.Update(tc, uLogin)
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, errors.New("refresh login affect rows != 1")
	}
	return uLogin, nil
}

func (l *login) findGroupUserInfos(c *tlog.Ctx, uid int64) ([]*vo.GroupUserInfo, error) {
	gUsers, err := Group.FindGroupUsersByUid(c, uid)
	if err != nil {
		c.Errorf("find group users by uid error: %v", err)
		return nil, err
	}
	if len(gUsers) == 0 {
		return nil, nil
	}
	var gInfos []*vo.GroupUserInfo
	for _, gu := range gUsers {
		gInfos = append(gInfos, gu.ToInfo())
	}
	return gInfos, nil
}
