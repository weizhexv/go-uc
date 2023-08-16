package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"errors"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/typeutil"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/var/roles"
)

type group struct {
}

var Group = new(group)

func (g *group) CreateGroup(c *tlog.Ctx, grp *dal.Group) (int64, error) {
	c.Infof("create group params: %v", grp)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("new trans error: %v", err)
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	rows, err := repo.Group.Insert(tc, grp)
	if err != nil {
		c.Errorf("group insert error: %v", err)
		return 0, err
	}

	c.Infof("group insert rows: %d", rows)
	return grp.Id, nil
}

func (g *group) UpdateGroup(c *tlog.Ctx, grp *dal.Group) (int64, error) {
	c.Infof("update group by %v", grp)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("new trans context error: %v", err)
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	rows, err := repo.Group.Update(tc, grp)
	if err != nil {
		c.Errorf("update group error: %v", err)
		return 0, err
	}
	c.Infof("update group rows: %d", rows)
	return grp.Id, nil
}

func (g *group) FindById(c *tlog.Ctx, id int64) (*dal.Group, error) {
	c.Infof("find group by id: %d", id)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	grp, err := repo.Group.GetById(tc, id)
	if err != nil {
		c.Errorf("find group by id error: %v", err)
		return nil, err
	}
	c.Infof("find group by id ret: %v", grp)
	return grp, nil
}

func (g *group) FindGroups(c *tlog.Ctx, domain domains.Domain, domainId int64, deleted any) ([]*dal.Group, error) {
	c.Infof("find groups by domain %v, domainId %d, deleted %v", domain, domainId, deleted)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find groups new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	ret, err := g.findGroups(c, tc, domain, domainId, deleted)
	if err != nil {
		c.Errorf("find groups error: %v", err)
	}
	return ret, err
}

func (g *group) findGroups(c *tlog.Ctx, tc *daog.TransContext, domain domains.Domain, domainId int64, deleted any) ([]*dal.Group, error) {
	c.Infof("find groups by domain %v, domainId %d, deleted %v", domain, domainId, deleted)
	mc := daog.NewMatcher().Eq(dal.GroupFields.Domain, string(domain))
	mc.Eq(dal.GroupFields.DomainId, domainId)
	deleted, checkDeleted := typeutil.GetBool(deleted)
	if checkDeleted {
		mc.Eq(dal.GroupFields.Deleted, deleted)
	}

	ret, err := repo.Group.QueryListMatcher(tc, mc, daog.NewOrder(dal.GroupFields.Deleted), daog.NewDescOrder(dal.GroupFields.CreatedAt))
	if err != nil {
		c.Errorf("find groups error: %v", err)
		return nil, err
	}

	return ret, nil
}

func (g *group) AddGroupUser(c *tlog.Ctx, tc *daog.TransContext, gu *dal.GroupUser) (int64, error) {
	c.Infof("add group user by: %v", gu)
	exist, err := repo.GroupUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.GroupUserFields.GroupId, gu.GroupId).Eq(dal.GroupUserFields.Uid, gu.Uid))
	if err != nil {
		c.Errorf("query group user error: %v", err)
		return 0, err
	}
	if exist != nil {
		return 0, errs.UserAlreadyInGroup
	}

	rows, err := repo.GroupUser.Insert(tc, gu)
	if err != nil {
		c.Errorf("insert group user error: %v", err)
		return 0, err
	}
	c.Infof("add group user result rows: %d", rows)
	return gu.Id, nil
}

func (g *group) RemoveGroupUser(c *tlog.Ctx, groupId int64, uid int64) error {
	c.Infof("remove group user by groupId %d, uid %d", groupId, uid)

	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("remove group user new trans error: %v", err)
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	p, err := repo.GroupUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.GroupUserFields.GroupId, groupId).Eq(dal.GroupUserFields.Uid, uid))
	if err != nil {
		c.Errorf("query remove group error: %v", err)
		return err
	}
	if p == nil {
		c.Infof("group not exist")
		return nil
	}

	rows, err := repo.GroupUser.DeleteById(tc, p.Id)
	c.Infof("delete group by id rows: %d", rows)

	if err != nil {
		c.Errorf("delete group by id error: %v", err)
	}

	return err
}

func (g *group) FindGroupUser(c *tlog.Ctx, groupId int64, uid int64) (*dal.GroupUser, error) {
	c.Infof("find group user by groupId %d, uid %d", groupId, uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find group user new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	gu, err := repo.GroupUser.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.GroupUserFields.GroupId, groupId).Eq(dal.GroupUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find group user error: %v", err)
		return nil, err
	}

	c.Infof("find group user result: %v", gu)
	return gu, nil
}

func (g *group) UpdateGroupUserRole(c *tlog.Ctx, id int64, role roles.Role) error {
	c.Infof("update group user role by id %d, role %v", id, role)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("update group user new trans error: %v", err)
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	rows, err := repo.GroupUser.UpdateById(tc, daog.NewModifier().Add(dal.GroupUserFields.RoleId, role.Id), id)
	c.Infof("update group user role result rows: %d", rows)

	if err != nil {
		c.Errorf("update group user role by id error: %v", err)
	}
	return err
}

func (g *group) FindGroupUsers(c *tlog.Ctx, groupId int64, uids []int64) ([]*dal.GroupUser, error) {
	c.Infof("find group users by groupId %d, uids %v", groupId, uids)

	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find group new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	mc := daog.NewMatcher().Eq(dal.GroupUserFields.GroupId, groupId)
	if len(uids) != 0 {
		mc.In(dal.GroupUserFields.Uid, daog.ConvertToAnySlice(uids))
	}

	ret, err := repo.GroupUser.QueryListMatcher(tc, mc)
	if err != nil {
		c.Errorf("find group users error: %v", err)
		return nil, err
	}
	return ret, nil
}

func (g *group) FindGroupUsersByUid(c *tlog.Ctx, uid int64) ([]*dal.GroupUser, error) {
	c.Infof("find group users by uid %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find group users new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	ret, err := repo.GroupUser.QueryListMatcher(tc, daog.NewMatcher().Eq(dal.GroupUserFields.Uid, uid))
	if err != nil {
		c.Errorf("find group users error: %v", err)
		return nil, err
	}
	return ret, nil
}

func (g *group) createDefault(c *tlog.Ctx, tc *daog.TransContext, domain domains.Domain, domainId int64) error {
	c.Infof("create default group by domain: %v, domainId %d", domain, domainId)
	groups, err := g.findGroups(c, tc, domain, domainId, nil)
	if err != nil {
		c.Errorf("find groups error: %v", err)
		return err
	}
	if len(groups) > 0 {
		c.Infof("group already exist")
		return nil
	}
	rows, err := repo.Group.Insert(tc, dal.NewDefaultGroup(domain, domainId))
	c.Infof("insert default group result: %d", rows)

	if err != nil {
		c.Errorf("insert default group error: %v", err)
	}
	return err
}

func (g *group) MapUidGroupRole(c *tlog.Ctx, tc *daog.TransContext, groupId int64, groupRole *roles.Role) (map[int64]*roles.Role, error) {
	c.Infof("map uid group role by groupId: %d, groupRole: %v", groupId, groupRole)
	if groupId <= 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}

	mc := daog.NewMatcher()
	mc.Eq(dal.GroupUserFields.GroupId, groupId)
	if groupRole != nil {
		mc.Eq(dal.GroupUserFields.RoleId, groupRole.Id)
	}

	ret, err := repo.GroupUser.QueryListMatcher(tc, mc)
	if err != nil {
		c.Errorf("map uid group role error: %v", err)
		return nil, err
	}

	mp := make(map[int64]*roles.Role)
	if len(ret) > 0 {
		for _, e := range ret {
			mp[e.Uid] = roles.OfId(e.RoleId)
		}
	}

	c.Infof("map uid group role ret: %v", mp)
	return mp, nil
}
