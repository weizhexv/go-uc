package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"github.com/rolandhe/daog"
	"go-uc/internal/convs"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/maputil"
	"go-uc/internal/tool/times"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

var BusinessUserImpl = new(businessUser)

type businessUser struct{}

func (b *businessUser) domain() domains.Domain {
	return domains.Business
}

func (b *businessUser) Query(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) ([]*vo.UserInfo, error) {
	c.Infof("business user query: %v", query)

	var uidGroupRoles map[int64]*roles.Role
	if query.GroupId > 0 {
		mp, err := Group.MapUidGroupRole(c, tc, query.GroupId, query.GroupRole)
		if err != nil {
			c.Errorf("group map uid group role error: %v", err)
			return nil, err
		}
		if len(mp) == 0 {
			return []*vo.UserInfo{}, nil
		}
		uidGroupRoles = mp
	}
	mc, err := b.queryToMatcher(query, maputil.CollectKeysOfMap(uidGroupRoles))
	if err != nil {
		c.Errorf("query to matcher error: %v", err)
		return nil, err
	}
	page, err := repo.BusinessUser.QueryPageListMatcher(tc, mc, query.ToPager())
	if err != nil {
		c.Errorf("business user page error: %v", err)
		return nil, err
	}
	return b.toInfos(page, uidGroupRoles), nil
}

func (b *businessUser) toInfos(pos []*dal.BusinessUser, uidGroupRoles map[int64]*roles.Role) []*vo.UserInfo {
	if len(pos) == 0 {
		return []*vo.UserInfo{}
	}
	sl := make([]*vo.UserInfo, len(pos))
	i := 0
	for _, e := range pos {
		sl[i] = e.ToInfo(uidGroupRoles)
		i++
	}
	return sl
}

func (b *businessUser) queryToMatcher(query *model.UserQuery, uids []int64) (daog.Matcher, error) {
	mc := daog.NewMatcher()
	if query.DomainId > 0 {
		mc.Eq(dal.BusinessUserFields.CustomerId, query.DomainId)
	}
	if len(query.Name) > 0 {
		mc.Like(dal.BusinessUserFields.Name, query.Name, daog.LikeStyleAll)
	}
	if len(query.Mobile) > 0 {
		mc.Like(dal.BusinessUserFields.Mobile, query.Mobile, daog.LikeStyleAll)
	}
	if query.CheckEnabled {
		mc.Eq(dal.BusinessUserFields.Enabled, query.Enabled)
	}
	if len(query.Roles) > 0 {
		var roleIds []int
		for _, r := range query.Roles {
			roleIds = append(roleIds, r.Id)
		}
		mc.In(dal.BusinessUserFields.RoleId, daog.ConvertToAnySlice(roleIds))
	}
	if query.GroupId > 0 {
		mc.In(dal.BusinessUserFields.Uid, daog.ConvertToAnySlice(uids))
	}
	return mc, nil
}

func (b *businessUser) Count(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) (int64, error) {
	c.Infof("business user count by %v", query)

	var uids []int64
	if query.GroupId > 0 {
		mp, err := Group.MapUidGroupRole(c, tc, query.GroupId, query.GroupRole)
		if err != nil {
			c.Errorf("group map uid group role error: %v", err)
			return 0, err
		}
		uids = maputil.CollectKeysOfMap(mp)
		if len(uids) == 0 {
			return 0, nil
		}
	}
	mc, err := b.queryToMatcher(query, uids)
	if err != nil {
		c.Errorf("query to matcher error: %v", err)
		return 0, err
	}
	ct, err := repo.BusinessUser.Count(tc, mc)
	if err != nil {
		c.Errorf("business user count error: %v", err)
		return 0, err
	}
	return ct, nil
}

func (b *businessUser) Insert(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("business user insert by %v", upsert)

	po := convs.NewBusinessUserConv(upsert).ToPO(c.DgContext)
	rows, err := repo.BusinessUser.Insert(tc, po)
	c.Infof("business user insert rows: %d", rows)
	if err != nil {
		c.Errorf("business user insert error: %v", err)
		return 0, err
	}
	if upsert.GroupId > 0 && upsert.GroupRole != nil {
		err = b.addGroupUser(c, tc, upsert)
		if err != nil {
			c.Errorf("add group user error: %v", err)
			return 0, err
		}
	}

	err = Email.SendRegisterEmail(c, tc, upsert.Uid)
	if err != nil {
		c.Errorf("send register email error: %v", err)
		return 0, err
	}
	return po.Id, nil
}

func (b *businessUser) addGroupUser(ctx *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) error {
	gu, err := b.toAddGroupUser(ctx, upsert)
	if err != nil {
		return err
	}
	guId, err := Group.AddGroupUser(ctx, tc, gu)
	if err != nil {
		return err
	}
	ctx.Infof("group user id is %d", guId)
	return nil
}

func (b *businessUser) Update(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("business user update by %v", upsert)

	po, err := repo.BusinessUser.FindByUid(tc, upsert.Uid)
	if err != nil {
		c.Errorf("business user find by uid error: %v", err)
		return 0, err
	}
	if po == nil {
		return 0, dgerr.USER_NOT_EXISTS
	}

	err = convs.NewBusinessUserConv(upsert).UpdatePO(c.DgContext, po)
	if err != nil {
		c.Errorf("business convert to update po error: %v", err)
		return 0, err
	}

	rows, err := repo.BusinessUser.Update(tc, po)
	c.Infof("business user update rows: %d", rows)
	if err != nil {
		c.Errorf("business user update error: %v", err)
		return 0, err
	}

	//update group user
	if upsert.GroupId > 0 && upsert.GroupRole != nil && roles.IsNotSuperManager(roles.OfId(po.RoleId)) {
		c.Infof("need update group user: groupId %d, groupRole %v, role %v",
			upsert.GroupId, upsert.GroupRole, roles.OfId(po.RoleId))

		err = b.updateGroupUser(c, tc, upsert)
		if err != nil {
			c.Errorf("update group user error: %v", err)
			return 0, err
		}
	}

	return po.Uid, nil
}

func (b *businessUser) updateGroupUser(ctx *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) error {
	gu, err := b.toUpdateGroupUser(ctx, tc, upsert)
	if err != nil {
		return err
	}
	_, err = repo.GroupUser.Update(tc, gu)
	if err != nil {
		return err
	}
	return nil
}

func (b *businessUser) FindByEmail(c *tlog.Ctx, tc *daog.TransContext, email string) (*vo.UserInfo, error) {
	c.Infof("business user find by email, %s", email)

	po, err := repo.BusinessUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.BusinessUserFields.Email, email))
	if err != nil {
		c.Errorf("business user query one error: %v", err)
		return nil, err
	}
	if po == nil {
		return nil, nil
	}
	return po.ToInfo(nil), nil
}

func (b *businessUser) FindByUid(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserInfo, error) {
	c.Infof("business user find by uid: %d", uid)

	po, err := repo.BusinessUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.BusinessUserFields.Uid, uid))
	if err != nil {
		c.Errorf("business user query one error: %v", err)
		return nil, err
	}
	if po == nil {
		return nil, nil
	}
	return po.ToInfo(nil), nil
}

func (b *businessUser) FindByUids(c *tlog.Ctx, tc *daog.TransContext, uids []int64) ([]*vo.UserInfo, error) {
	c.Infof("business user find by uids: %v", uids)

	pos, err := repo.BusinessUser.QueryListMatcher(tc, daog.NewMatcher().In(dal.BusinessUserFields.Uid, daog.ConvertToAnySlice(uids)))
	if err != nil {
		c.Errorf("business user query list error: %v", err)
		return nil, err
	}

	return b.toInfos(pos, nil), nil
}

func (b *businessUser) ResetPassword(c *tlog.Ctx, tc *daog.TransContext, hash string, uid int64) error {
	c.Infof("business user reset password by hash:%s, uid:%d", hash, uid)

	usr, err := b.FindByUid(c, tc, uid)
	if err != nil {
		c.Errorf("business user find by uid error: %v", err)
		return err
	}
	if usr == nil {
		return dgerr.USER_NOT_EXISTS
	}
	rows, err := repo.BusinessUser.UpdateByModifier(tc, daog.NewModifier().Add(dal.BusinessUserFields.Password, hash),
		daog.NewMatcher().Eq(dal.BusinessUserFields.Uid, uid))
	if err != nil {
		c.Errorf("update business user error: %v", err)
		return err
	}

	c.Infof("update password success ret rows: %d", rows)
	return nil
}

func (b *businessUser) FindByDomainId(c *tlog.Ctx, tc *daog.TransContext, domainId int64) ([]*vo.UserInfo, error) {
	c.Infof("business user find by domainId %d", domainId)

	mc := daog.NewMatcher()
	if domainId > 0 {
		mc.Eq(dal.BusinessUserFields.CustomerId, domainId)
	}
	pos, err := repo.BusinessUser.QueryListMatcher(tc, mc)
	if err != nil {
		c.Errorf("business user query list error: %v", err)
		return nil, err
	}
	return b.toInfos(pos, nil), nil
}

func (b *businessUser) MapDomainHasSuperManager(c *tlog.Ctx, tc *daog.TransContext, domainIds []int64) (map[int64]bool, error) {
	c.Errorf("MapDomainHasSuperManager method not support!")
	return nil, dgerr.ILLEGAL_OPERATION
}

func (b *businessUser) FindUserModel(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*model.UserModel, error) {
	c.Infof("business user find user model by uid %d", uid)

	po, err := repo.BusinessUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.BusinessUserFields.Uid, uid))
	if err != nil {
		c.Infof("business user query one error: %v", err)
		return nil, err
	}
	if po == nil {
		return nil, nil
	}

	return po.ToModel(), nil
}

func (b *businessUser) toAddGroupUser(ctx *tlog.Ctx, upsert *model.UpsertUser) (*dal.GroupUser, error) {
	grp, err := Group.FindById(ctx, upsert.GroupId)
	if err != nil {
		return nil, err
	}
	if grp == nil {
		ctx.Errorf("group not exists, id %d", upsert.GroupId)
		return nil, errs.GroupNotExist
	}
	if domains.NotEqual(grp.Domain, domains.Business) {
		ctx.Errorf("group domain not equals %s != %s", grp.Domain, domains.Business)
		return nil, errs.GroupNotExist
	}
	if grp.DomainId != ctx.DomainId {
		ctx.Errorf("group domainId not equals %d != %d", grp.DomainId, ctx.DomainId)
		return nil, errs.GroupNotExist
	}
	return dal.NewGroupUser(ctx, upsert), nil
}

func (b *businessUser) toUpdateGroupUser(ctx *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (*dal.GroupUser, error) {
	grp, err := repo.GroupUser.FindByGroupIdAndUid(tc, upsert.GroupId, upsert.Uid)
	if err != nil {
		return nil, err
	}
	if grp == nil {
		return nil, errs.GroupUserNotExist
	}
	grp.RoleId = upsert.GroupRole.Id
	grp.ModifiedBy = ctx.UserId
	grp.ModifiedAt = times.NowDatetime()
	return grp, nil
}
