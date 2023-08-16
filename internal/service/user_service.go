package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"dghire.com/libs/go-common/page"
	"errors"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/passwords"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

type IUser interface {
	domain() domains.Domain
	Query(ctx *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) ([]*vo.UserInfo, error)
	Count(ctx *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) (int64, error)
	Insert(ctx *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error)
	Update(ctx *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error)
	FindByEmail(ctx *tlog.Ctx, tc *daog.TransContext, email string) (*vo.UserInfo, error)
	FindByUid(ctx *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserInfo, error)
	FindByUids(ctx *tlog.Ctx, tc *daog.TransContext, uids []int64) ([]*vo.UserInfo, error)
	ResetPassword(ctx *tlog.Ctx, tc *daog.TransContext, hash string, uid int64) error
	FindByDomainId(ctx *tlog.Ctx, tc *daog.TransContext, domainId int64) ([]*vo.UserInfo, error)
	MapDomainHasSuperManager(ctx *tlog.Ctx, tc *daog.TransContext, domainIds []int64) (map[int64]bool, error)
	FindUserModel(ctx *tlog.Ctx, tc *daog.TransContext, uid int64) (*model.UserModel, error)
}

var UserImpls = initUserImpls()

func initUserImpls() *userImpls {
	u := newUserImpls()
	u.holder[domains.Business] = BusinessUserImpl
	u.holder[domains.Platform] = PlatformUserImpl
	u.holder[domains.Supplier] = SupplierUserImpl
	u.holder[domains.Employee] = EmployeeImpl
	return u
}

type userImpls struct {
	holder map[domains.Domain]IUser
}

func newUserImpls() *userImpls {
	return &userImpls{
		holder: make(map[domains.Domain]IUser),
	}
}

func targetImpl(domain domains.Domain) (IUser, error) {
	impl := UserImpls.holder[domain]
	if impl == nil {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}
	return impl, nil
}

func (u *userImpls) Page(c *tlog.Ctx, query *model.UserQuery) (*page.PageList[vo.UserDetail], error) {
	c.Infof("page user by %v", query)
	infos, err := u.Query(c, query)
	if err != nil {
		c.Errorf("query user error: %v", err)
		return nil, err
	}
	userDetails, err := buildUserDetails(c, infos)
	if err != nil {
		c.Errorf("build user details error: %v", err)
		return nil, err
	}
	count, err := u.Count(c, query)
	if err != nil {
		c.Errorf("count user error: %v", err)
		return nil, err
	}
	return page.ListOf[vo.UserDetail](query.PageNo, query.PageSize, int(count), userDetails), nil
}

func buildUserDetails(c *tlog.Ctx, infos []*vo.UserInfo) ([]*vo.UserDetail, error) {
	if len(infos) == 0 {
		return []*vo.UserDetail{}, nil
	}
	var domainIds []int64
	for _, e := range infos {
		domainIds = append(domainIds, e.DomainId)
	}

	domainNames, err := DomainService.MapDomainIdNames(c, infos[0].Domain, domainIds)
	if err != nil {
		return nil, err
	}
	return vo.UserDetailsOf(c.DgContext, infos, domainNames), nil
}

func (u *userImpls) Query(c *tlog.Ctx, query *model.UserQuery) ([]*vo.UserInfo, error) {
	c.Infof("query user by %v", query)
	impl, err := targetImpl(query.Domain)
	if err != nil {
		c.Errorf("find target impl error: %v", err)
		return nil, err
	}
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("query user new trans error: %v", err)
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
	ret, err := impl.Query(c, tc, query)
	if err != nil {
		c.Errorf("query user error: %v", err)
	}
	return ret, err
}

func (u *userImpls) Count(c *tlog.Ctx, query *model.UserQuery) (int64, error) {
	c.Infof("count user by %v", query)
	impl, err := targetImpl(query.Domain)
	if err != nil {
		c.Errorf("find target impl error: %v", err)
		return 0, err
	}
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("count user new trans error: %v", err)
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
	ret, err := impl.Count(c, tc, query)
	if err != nil {
		c.Errorf("count user error: %v", err)
	}
	return ret, err
}

func (u *userImpls) Upsert(c *tlog.Ctx, upsert *model.UpsertUser) (int64, error) {
	if upsert.Insert {
		return u.insert(c, upsert)
	} else {
		return u.update(c, upsert)
	}
}

func (u *userImpls) insert(c *tlog.Ctx, upsert *model.UpsertUser) (int64, error) {
	c.Infof("insert user by %v", upsert)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("insert user new trans error: %v", err)
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

	//Insert index
	idx := dal.NewUserIndex(c, upsert.Domain, upsert.Email)
	rows, err := repo.UserIndex.Insert(tc, idx)
	c.Infof("insert user index rows %d uid %d", rows, idx.Id)
	if err != nil {
		c.Errorf("insert user index error: %v", err)
		return 0, err
	}

	//Insert user
	impl, err := targetImpl(upsert.Domain)
	if err != nil {
		c.Errorf("find target impl error: %v", err)
		return 0, err
	}
	upsert.Uid = idx.Id
	subId, err := impl.Insert(c, tc, upsert)
	if err != nil {
		c.Errorf("impl insert error: %v", err)
		return 0, err
	}
	c.Infof("insert user subId %d", subId)

	//Update index.subId
	rows, err = repo.UserIndex.UpdateById(tc, daog.NewModifier().Add(dal.UserIndexFields.SubId, subId), idx.Id)
	if err != nil {
		c.Errorf("update user index error: %v", err)
		return 0, err
	}
	c.Infof("update user index rows: %d", rows)
	return idx.Id, err
}

func (u *userImpls) update(c *tlog.Ctx, upsert *model.UpsertUser) (int64, error) {
	c.Infof("update user by %v", upsert)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Infof("update user new trans error: %v", err)
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

	//check index exists
	index, err := repo.UserIndex.GetById(tc, upsert.Uid)
	if err != nil {
		c.Error("find user index error:", err)
		return 0, err
	}
	if index == nil {
		return 0, dgerr.USER_NOT_EXISTS
	}

	//Update index identity if changed
	if len(upsert.Email) != 0 {
		_, err = repo.UserIndex.UpdateById(tc, daog.NewModifier().Add(dal.UserIndexFields.Identity, upsert.Email), upsert.Uid)
		if err != nil {
			c.Errorf("update user index error: %v", err)
			return 0, err
		}
	}

	//Update user
	upsert.Id = index.SubId
	impl, err := targetImpl(upsert.Domain)
	if err != nil {
		c.Errorf("find target impl error: %v", err)
		return 0, err
	}
	id, err := impl.Update(c, tc, upsert)
	if err != nil {
		c.Errorf("impl update error: %v", err)
	}
	return id, err
}

func (u *userImpls) FindByEmail(c *tlog.Ctx, email string) (*vo.UserInfo, error) {
	c.Infof("find user by email: %s", email)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find user by email new trans error: %v", err)
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

	//find index first
	index, err := repo.UserIndex.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.UserIndexFields.Identity, email))
	if err != nil {
		c.Errorf("find user index error: %v", err)
		return nil, err
	}
	if index == nil {
		return nil, nil
	}

	//find domain user
	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		c.Errorf("find traget impl error: %v", err)
		return nil, err
	}
	ret, err := impl.FindByEmail(c, tc, email)
	if err != nil {
		c.Errorf("impl find by email error: %v", err)
	}
	return ret, err
}

func (u *userImpls) FindByUid(c *tlog.Ctx, uid int64) (*vo.UserInfo, error) {
	c.Infof("find user by uid %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find user by uid new trans error: %v", err)
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

	c.Infof("find user index by uid: %d", uid)
	index, err := repo.UserIndex.GetById(tc, uid)
	if err != nil || index == nil {
		return nil, dgerr.USER_NOT_EXISTS
	}

	//find domain user
	c.Infof("find target by domain | %s", index.Domain)
	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		c.Errorf("find target impl error: %v", err)
		return nil, err
	}
	c.Infof("find impl user by uid | %d", uid)
	ret, err := impl.FindByUid(c, tc, uid)
	if err != nil {
		c.Errorf("impl find by uid error: %v", err)
	}
	return ret, err
}

func (u *userImpls) FindByUids(c *tlog.Ctx, uids []int64) ([]*vo.UserInfo, error) {
	c.Infof("find user by uids: %v", uids)
	if len(uids) == 0 {
		return []*vo.UserInfo{}, nil
	}

	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find user by uids new trans error: %v", err)
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

	indexes, err := repo.UserIndex.GetByIds(tc, uids)
	if err != nil {
		c.Errorf("find user indexes error: %v", err)
		return nil, err
	}

	mp := make(map[domains.Domain]bool)
	for _, inx := range indexes {
		dm := domains.Parse(inx.Domain)
		if len(dm) != 0 {
			mp[dm] = true
		}
	}

	var infos []*vo.UserInfo
	for k, _ := range mp {
		impl, err := targetImpl(k)
		if err != nil {
			c.Error("not find target user impl:", k)
			continue
		}
		ret, err := impl.FindByUids(c, tc, uids)
		if err != nil {
			c.Errorf("impl find by uids error: %v", err)
			return nil, err
		}
		infos = append(infos, ret...)
	}
	return infos, nil
}

func (u *userImpls) ResetPassword(c *tlog.Ctx, pwd string, uid int64) error {
	c.Infof("reset password by pwd %s, uid %d", pwd, uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
		c.Errorf("reset pwd new trans error: %v", err)
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

	index, err := repo.UserIndex.GetById(tc, uid)
	if err != nil || index == nil {
		return dgerr.USER_NOT_EXISTS
	}

	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		return err
	}
	hash, err := passwords.Hash(pwd)
	if err != nil {
		c.Errorf("password hash error: %v", err)
		return err
	}
	err = impl.ResetPassword(c, tc, hash, uid)
	if err != nil {
		c.Errorf("reset password error: %v", err)
	}
	return err
}

func (u *userImpls) FindByDomainId(c *tlog.Ctx, domain domains.Domain, domainId int64) ([]*vo.UserInfo, error) {
	c.Infof("find user by domainId, domain: %v, domainId: %d", domain, domainId)
	impl, err := targetImpl(domain)
	if err != nil {
		return nil, err
	}

	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find user by domainId new trans error: %v", err)
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

	ret, err := impl.FindByDomainId(c, tc, domainId)
	if err != nil {
		c.Errorf("find user by domainId error: %v", err)
	}
	return ret, err
}

func (u *userImpls) mapDomainHasSuperManager(c *tlog.Ctx, domain domains.Domain, domainIds []int64) (map[int64]bool, error) {
	c.Infof("map domain has super manager by domain %v, domainIds %v", domain, domainIds)
	impl, err := targetImpl(domain)
	if err != nil {
		return nil, err
	}

	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
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

	ret, err := impl.MapDomainHasSuperManager(c, tc, domainIds)
	if err != nil {
		c.Errorf("map domain has super manager error: %v", err)
	}
	return ret, err
}

func (u *userImpls) FindByDomainAndRole(c *tlog.Ctx, domain domains.Domain, domainId int64, role *roles.Role) ([]*vo.UserInfo, error) {
	c.Infof("find by domain and role, domain %v, domainId %d, role %v", domain, domainId, role)

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

	impl, err := targetImpl(domain)
	if err != nil {
		return nil, err
	}
	ret, err := impl.Query(c, tc, &model.UserQuery{
		DomainId: domainId,
		Roles:    []*roles.Role{role},
	})
	if err != nil {
		c.Errorf("impl query error: %v", err)
	}
	return ret, err
}

func (u *userImpls) FindSeedByUid(c *tlog.Ctx, uid int64) ([]byte, error) {
	c.Infof("find seed by uid %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
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

	index, err := repo.UserIndex.GetById(tc, uid)
	if err != nil || index == nil {
		return nil, dgerr.USER_NOT_EXISTS
	}

	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		return nil, err
	}
	userPo, err := impl.FindUserModel(c, tc, uid)
	if err != nil {
		return nil, err
	}
	return userPo.Seed, nil
}

func (u *userImpls) FindPwdByUid(c *tlog.Ctx, uid int64) (string, error) {
	c.Infof("find pwd by uid %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		return "", err
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

	index, err := repo.UserIndex.GetById(tc, uid)
	if err != nil || index == nil {
		return "", dgerr.USER_NOT_EXISTS
	}

	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		return "", err
	}
	userPo, err := impl.FindUserModel(c, tc, uid)
	if err != nil {
		return "", err
	}
	return userPo.Password, nil
}

func (u *userImpls) FindDetailByUid(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserDetail, error) {
	c.Infof("find detail by uid: %d", uid)
	index, err := repo.UserIndex.GetById(tc, uid)
	if err != nil || index == nil {
		return nil, dgerr.USER_NOT_EXISTS
	}

	impl, err := targetImpl(domains.Parse(index.Domain))
	if err != nil {
		return nil, err
	}
	info, err := impl.FindByUid(c, tc, uid)
	if err != nil {
		return nil, err
	}
	domainName := ""
	if info.Domain.Is(domains.Business) {
		c, err := Customer.FindById(c, info.DomainId)
		if err != nil {
			return nil, err
		}
		if c != nil {
			domainName = c.FullName
		}
	} else if info.Domain.Is(domains.Supplier) {
		s, err := Supplier.FindById(c, info.DomainId)
		if err != nil {
			return nil, err
		}
		if s != nil {
			domainName = s.FullName
		}
	}

	return vo.UserDetailOf(c.DgContext, info, domainName), nil
}
