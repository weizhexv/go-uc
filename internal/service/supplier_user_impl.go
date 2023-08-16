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
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

var SupplierUserImpl = new(supplierUser)

type supplierUser struct{}

func (s *supplierUser) domain() domains.Domain {
	return domains.Supplier
}

func (s *supplierUser) Query(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) ([]*vo.UserInfo, error) {
	c.Infof("query supplier user by %v", query)
	ret, err := repo.SupplierUser.QueryPageListMatcher(tc, s.queryToMatcher(query), query.ToPager())
	if err != nil {
		c.Errorf("query supplier user error: %v", err)
		return nil, err
	}
	return s.toInfos(ret), nil
}

func (s *supplierUser) toInfos(users []*dal.SupplierUser) []*vo.UserInfo {
	var infos []*vo.UserInfo
	if len(users) > 0 {
		for _, u := range users {
			infos = append(infos, u.ToInfo())
		}
	}
	return infos
}

func (s *supplierUser) queryToMatcher(query *model.UserQuery) daog.Matcher {
	mc := daog.NewMatcher()
	if query.DomainId > 0 {
		mc.Eq(dal.SupplierUserFields.SupplierId, query.DomainId)
	}
	if len(query.Name) > 0 {
		mc.Like(dal.SupplierUserFields.Name, query.Name, daog.LikeStyleAll)
	}
	if len(query.Mobile) > 0 {
		mc.Like(dal.SupplierUserFields.Mobile, query.Mobile, daog.LikeStyleAll)
	}
	if len(query.Email) > 0 {
		mc.Like(dal.SupplierUserFields.Email, query.Email, daog.LikeStyleAll)
	}
	if query.CheckEnabled {
		mc.Eq(dal.SupplierUserFields.Enabled, query.Enabled)
	}
	if len(query.Roles) > 0 {
		var roleIds []int
		for _, r := range query.Roles {
			roleIds = append(roleIds, r.Id)
		}
		mc.In(dal.SupplierUserFields.RoleId, daog.ConvertToAnySlice(roleIds))
	}
	return mc
}

func (s *supplierUser) Count(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) (int64, error) {
	c.Infof("count supplier user by %v", query)
	ct, err := repo.SupplierUser.Count(tc, s.queryToMatcher(query))
	if err != nil {
		c.Errorf("count supplier user error: %v", err)
		return 0, err
	}
	c.Infof("count supplier user ret: %d", ct)
	return ct, nil
}

func (s *supplierUser) Insert(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("insert supplier user by %v", upsert)
	po := convs.NewSupplierUserConv(upsert).ToPO(c.DgContext)
	rows, err := repo.SupplierUser.Insert(tc, po)
	if err != nil {
		c.Errorf("insert supplier user error: %v", err)
		return 0, err
	}
	c.Infof("insert supplier user ret rows: %d", rows)
	err = Email.SendRegisterEmail(c, tc, upsert.Uid)
	if err != nil {
		c.Errorf("send register email error: %v", err)
		return 0, err
	}
	return po.Id, nil
}

func (s *supplierUser) Update(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("supplier user update by %v", upsert)
	po, err := repo.SupplierUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.SupplierUserFields.Id, upsert.Id))
	if err != nil {
		c.Errorf("query supplier user error: %v", err)
		return 0, nil
	}
	if po == nil {
		c.Errorf("supplier user not exist")
		return 0, dgerr.USER_NOT_EXISTS
	}
	err = convs.NewSupplierUserConv(upsert).UpdatePO(c.DgContext, po)
	if err != nil {
		c.Errorf("supplier user convs error: %v", err)
		return 0, err
	}
	rows, err := repo.SupplierUser.Update(tc, po)
	if err != nil {
		c.Errorf("supplier user update error: %v", err)
		return 0, err
	}
	c.Infof("supplier user update affect rows: %d", rows)
	return po.Uid, nil
}

func (s *supplierUser) FindByEmail(c *tlog.Ctx, tc *daog.TransContext, email string) (*vo.UserInfo, error) {
	c.Infof("find supplier user by email: %s", email)
	if len(email) == 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}
	ret, err := repo.SupplierUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.SupplierUserFields.Email, email))
	if err != nil {
		c.Errorf("find supplier user error: %v", err)
		return nil, err
	}
	return ret.ToInfo(), nil
}

func (s *supplierUser) FindByUid(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserInfo, error) {
	c.Infof("find supplier user by uid: %d", uid)
	if uid <= 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}
	ret, err := repo.SupplierUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.SupplierUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find supplier user error: %v", err)
		return nil, err
	}
	return ret.ToInfo(), nil
}

func (s *supplierUser) FindByUids(c *tlog.Ctx, tc *daog.TransContext, uids []int64) ([]*vo.UserInfo, error) {
	c.Infof("find supplier user by uids: %v", uids)
	if len(uids) == 0 {
		return []*vo.UserInfo{}, nil
	}
	ret, err := repo.SupplierUser.QueryListMatcher(tc, daog.NewMatcher().In(dal.SupplierUserFields.Uid, daog.ConvertToAnySlice(uids)))
	if err != nil {
		c.Errorf("find supplier user by uids error: %v", err)
		return nil, err
	}
	return s.toInfos(ret), nil
}

func (s *supplierUser) ResetPassword(c *tlog.Ctx, tc *daog.TransContext, hash string, uid int64) error {
	c.Infof("supplier user reset pwd by hash %s, uid %d", hash, uid)
	if uid <= 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	exist, err := s.FindByUid(c, tc, uid)
	if err != nil {
		c.Errorf("supplier user find by uid error: %v", err)
		return err
	}
	if exist == nil {
		c.Errorf("supplier user not exist")
		return dgerr.USER_NOT_EXISTS
	}
	rows, err := repo.SupplierUser.UpdateByModifier(tc, daog.NewModifier().Add(dal.SupplierUserFields.Password, hash),
		daog.NewMatcher().Eq(dal.SupplierUserFields.Uid, uid))
	if err != nil {
		c.Errorf("supplier user reset pwd error: %v", err)
		return err
	}
	c.Infof("supplier user reset pwd ret rows: %d", rows)
	return nil
}

func (s *supplierUser) FindByDomainId(c *tlog.Ctx, tc *daog.TransContext, domainId int64) ([]*vo.UserInfo, error) {
	c.Infof("supplier user find by domainId %d", domainId)
	mc := daog.NewMatcher()
	if domainId > 0 {
		mc.Eq(dal.SupplierUserFields.SupplierId, domainId)
	}
	ret, err := repo.SupplierUser.QueryListMatcher(tc, mc)
	if err != nil {
		c.Errorf("supplier user find by domainId error: %v", err)
		return nil, err
	}
	return s.toInfos(ret), err
}

func (s *supplierUser) MapDomainHasSuperManager(c *tlog.Ctx, tc *daog.TransContext, domainIds []int64) (map[int64]bool, error) {
	c.Infof("map domain has super manager by domainIds: %v", domainIds)
	mp := make(map[int64]bool)
	if len(domainIds) == 0 {
		return mp, nil
	}
	mc := daog.NewMatcher().In(dal.SupplierUserFields.SupplierId, daog.ConvertToAnySlice(domainIds))
	mc.Eq(dal.SupplierUserFields.RoleId, roles.SuperManager.Id)
	mc.Eq(dal.SupplierUserFields.Enabled, true)
	ret, err := repo.SupplierUser.QueryListMatcher(tc, mc)
	if err != nil {
		c.Errorf("query supplier user list error: %v", err)
		return nil, err
	}
	if len(ret) == 0 {
		return mp, nil
	}
	for _, domainId := range domainIds {
		mp[domainId] = false
	}
	for _, po := range ret {
		mp[po.SupplierId] = true
	}
	c.Infof("map domain has super manager ret map: %v", mp)
	return mp, nil
}

func (s *supplierUser) FindUserModel(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*model.UserModel, error) {
	c.Infof("find supplier user model by uid %d", uid)
	po, err := repo.SupplierUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.SupplierUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find supplier user model error: %v", err)
		return nil, err
	}
	return po.ToModel(), nil
}
