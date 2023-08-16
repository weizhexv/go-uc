package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"github.com/rolandhe/daog"
	"go-uc/internal/convs"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/vo"
)

var PlatformUserImpl = new(platformUser)

type platformUser struct{}

func (p *platformUser) domain() domains.Domain {
	return domains.Platform
}

func (p *platformUser) Query(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) ([]*vo.UserInfo, error) {
	c.Infof("query platform user by %v", query)
	ret, err := repo.PlatformUser.QueryPageListMatcher(tc, p.queryToMatcher(query), query.ToPager())
	if err != nil {
		c.Errorf("query platform user error: %v", err)
		return nil, err
	}
	return p.toInfos(ret), nil
}

func (p *platformUser) toInfos(users []*dal.PlatformUser) []*vo.UserInfo {
	var infos []*vo.UserInfo
	if len(users) > 0 {
		for _, u := range users {
			infos = append(infos, u.ToInfo())
		}
	}
	return infos
}

func (p *platformUser) queryToMatcher(query *model.UserQuery) daog.Matcher {
	mc := daog.NewMatcher()
	if len(query.Name) > 0 {
		mc.Like(dal.PlatformUserFields.Name, query.Name, daog.LikeStyleAll)
	}
	if len(query.Mobile) > 0 {
		mc.Like(dal.PlatformUserFields.Mobile, query.Mobile, daog.LikeStyleAll)
	}
	if len(query.Email) > 0 {
		mc.Like(dal.PlatformUserFields.Email, query.Email, daog.LikeStyleAll)
	}
	if query.CheckEnabled {
		mc.Eq(dal.PlatformUserFields.Enabled, query.Enabled)
	}
	if len(query.Roles) > 0 {
		var roleIds []int
		for _, r := range query.Roles {
			roleIds = append(roleIds, r.Id)
		}
		mc.In(dal.PlatformUserFields.RoleId, daog.ConvertToAnySlice(roleIds))
	}
	return mc
}

func (p *platformUser) Count(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) (int64, error) {
	c.Infof("count platform user by %v", query)
	ct, err := repo.PlatformUser.Count(tc, p.queryToMatcher(query))
	if err != nil {
		c.Errorf("count platform user error: %v", err)
		return 0, err
	}
	c.Infof("count platfom user ret: %d", ct)
	return ct, nil
}

func (p *platformUser) Insert(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("insert platform user by: %v", upsert)

	po := convs.NewPlatformUserConv(upsert).ToPO(c.DgContext)
	rows, err := repo.PlatformUser.Insert(tc, po)
	if err != nil {
		c.Errorf("insert platform user error: %v", err)
		return 0, err
	}
	c.Infof("insert platform user ret rows: %d", rows)

	err = Email.SendRegisterEmail(c, tc, upsert.Uid)
	if err != nil {
		c.Errorf("send platform register email error: %v", err)
		return 0, err
	}
	return po.Id, nil
}

func (p *platformUser) Update(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("update platform user by %v", upsert)
	po, err := p.findPO(c, tc, upsert.Uid)
	if err != nil {
		c.Errorf("find platform user po error: %v", err)
		return 0, err
	}
	err = convs.NewPlatformUserConv(upsert).UpdatePO(c.DgContext, po)
	if err != nil {
		c.Errorf("platform user convs error: %v", err)
		return 0, err
	}
	rows, err := repo.PlatformUser.Update(tc, po)
	if err != nil {
		c.Errorf("update platform user error: %v", err)
		return 0, err
	}

	c.Infof("update platform user ret rows: %d", rows)
	return po.Uid, nil
}

func (p *platformUser) FindByEmail(c *tlog.Ctx, tc *daog.TransContext, email string) (*vo.UserInfo, error) {
	c.Infof("find platform user by email: %s", email)
	if len(email) == 0 {
		c.Warnf("find platform user blank email param")
		return nil, nil
	}
	po, err := repo.PlatformUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.PlatformUserFields.Email, email))
	if err != nil {
		c.Errorf("platform user query one error: %v", err)
		return nil, err
	}
	return po.ToInfo(), err
}

func (p *platformUser) FindByUid(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserInfo, error) {
	c.Infof("find platform user by uid: %d", uid)
	if uid <= 0 {
		c.Warnf("find platform user uid param less or equal to 0")
		return nil, nil
	}
	po, err := repo.PlatformUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.PlatformUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find platform user error: %v", err)
		return nil, err
	}
	return po.ToInfo(), err
}

func (p *platformUser) FindByUids(c *tlog.Ctx, tc *daog.TransContext, uids []int64) ([]*vo.UserInfo, error) {
	c.Infof("find platform user by uids: %v", uids)
	if len(uids) == 0 {
		return []*vo.UserInfo{}, nil
	}
	users, err := repo.PlatformUser.QueryListMatcher(tc, daog.NewMatcher().In(dal.PlatformUserFields.Uid, daog.ConvertToAnySlice(uids)))
	if err != nil {
		c.Errorf("find list platform user error: %v", err)
		return nil, err
	}
	return p.toInfos(users), nil
}

func (p *platformUser) ResetPassword(c *tlog.Ctx, tc *daog.TransContext, hash string, uid int64) error {
	c.Infof("platform user reset password by hash: %s, uid: %d", hash, uid)
	if len(hash) == 0 || uid <= 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	info, err := p.FindByUid(c, tc, uid)
	if err != nil {
		c.Errorf("find platform user error: %v", err)
		return err
	}
	if info == nil {
		c.Errorf("platform user not exist")
		return dgerr.USER_NOT_EXISTS
	}
	rows, err := repo.PlatformUser.UpdateByModifier(tc, daog.NewModifier().Add(dal.PlatformUserFields.Password, hash),
		daog.NewMatcher().Eq(dal.PlatformUserFields.Uid, uid))
	if err != nil {
		c.Errorf("platform user reset password error: %v", err)
	}

	c.Infof("platform user reset password result rows: %d", rows)
	return err
}

func (p *platformUser) FindByDomainId(c *tlog.Ctx, tc *daog.TransContext, domainId int64) ([]*vo.UserInfo, error) {
	c.Infof("find all platform user")
	all, err := repo.PlatformUser.GetAll(tc)
	if err != nil {
		c.Errorf("find all platform user error: %v", err)
		return nil, err
	}
	return p.toInfos(all), nil
}

func (p *platformUser) MapDomainHasSuperManager(ctx *tlog.Ctx, tc *daog.TransContext, domainIds []int64) (map[int64]bool, error) {
	return nil, dgerr.ILLEGAL_OPERATION
}

func (p *platformUser) findPO(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*dal.PlatformUser, error) {
	c.Infof("find platform user po by uid: %d", uid)
	po, err := repo.PlatformUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.PlatformUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find platform user po error: %v", err)
		return nil, err
	}
	return po, nil
}

func (p *platformUser) FindUserModel(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*model.UserModel, error) {
	c.Infof("find platform user model by uid: %d", uid)
	po, err := repo.PlatformUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.PlatformUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find platform user model error: %v", err)
		return nil, err
	}
	return po.ToModel(), nil
}
