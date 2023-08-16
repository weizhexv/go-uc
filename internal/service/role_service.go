package service

import (
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"golang.org/x/exp/slices"
)

type role struct {
}

var Role = new(role)

func (r *role) GetRolePermissions(c tlog.Ctx, domain string, roleId int) ([]string, error) {
	c.Infof("get role permissions by domain %s, roleId %d", domain, roleId)

	return daog.AutoTransWithResult[[]string](func() (*daog.TransContext, error) {
		return daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	}, func(tc *daog.TransContext) ([]string, error) {
		rps, err := repo.RolePermission.QueryListMatcher(tc,
			daog.NewMatcher().Eq(dal.RolePermissionFields.Domain, domain).Eq(dal.RolePermissionFields.RoleId, roleId))
		if err != nil {
			c.Errorf("get role permissions error: %v", err)
			return nil, err
		}

		var sl []string
		if len(rps) > 0 {
			for _, o := range rps {
				sl = append(sl, o.Permission)
			}
		}
		c.Infof("get role permissions ret: %v", sl)
		return sl, nil
	})
}

func (r *role) GetIdNameMap(roleIds []int) map[int]string {
	ret := make(map[int]string)
	if len(roleIds) == 0 {
		return ret
	}
	for _, r := range roleIds {
		o := roles.OfId(r)
		if o != nil {
			ret[o.Id] = o.Name
		}
	}
	return ret
}

func (r *role) FindByDomain(domain domains.Domain) []*roles.Role {
	var ret []*roles.Role
	for _, r := range roles.All() {
		if slices.Contains(r.Domains, domain) {
			ret = append(ret, r)
		}
	}
	return ret
}

func (r *role) FindByDomainMap(domain domains.Domain) map[int]*roles.Role {
	ret := make(map[int]*roles.Role)
	rs := r.FindByDomain(domain)
	if len(rs) == 0 {
		return ret
	}
	for _, r := range rs {
		ret[r.Id] = r
	}
	return ret
}

func (r *role) ContainsRole(domain domains.Domain, role roles.Role) bool {
	mp := r.FindByDomainMap(domain)
	_, ok := mp[role.Id]
	return ok
}

func (r *role) HasSuperManager(c *tlog.Ctx, domain domains.Domain, domainId int64) (bool, error) {
	c.Infof("has super manager by domain: %v, domainId: %d", domain, domainId)

	if domain.IsAny(domains.Platform, domains.Employee) {
		return false, nil
	}
	q := &model.UserQuery{
		Domain:       domain,
		DomainId:     domainId,
		CheckEnabled: true,
		Enabled:      true,
		Roles:        []*roles.Role{roles.SuperManager},
		PageNo:       1,
		PageSize:     1,
	}

	c.Infof("query super manager by %v", q)
	ret, err := UserImpls.Query(c, q)
	if err != nil {
		c.Errorf("query super manager error: %v", err)
		return false, err
	}
	return len(ret) > 0, nil
}
