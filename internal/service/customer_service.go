package service

import (
	"errors"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
)

var Customer = new(customer)

type customer struct {
}

func (c *customer) Query(ctx *tlog.Ctx, query *model.CustomerQuery) ([]*dal.Customer, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	ret, err := repo.Customer.QueryListMatcher(tc, c.queryToMatch(query))
	return ret, err
}

func (c *customer) Count(ctx *tlog.Ctx, query *model.CustomerQuery) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	if err != nil {
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	ct, err := repo.Customer.Count(tc, c.queryToMatch(query))
	if err != nil {
		return 0, err
	}
	return ct, nil
}

func (c *customer) queryToMatch(query *model.CustomerQuery) daog.Matcher {
	mc := daog.NewMatcher()
	if query.Id > 0 {
		mc.Eq(dal.CustomerFields.Id, query.Id)
	}
	if len(query.FullName) > 0 {
		mc.Like(dal.CustomerFields.FullName, query.FullName, daog.LikeStyleAll)
	}
	if query.CheckEnabled {
		mc.Eq(dal.CustomerFields.Enabled, query.Enabled)
	}
	if len(query.ProjectId) > 0 {
		mc.Eq(dal.CustomerFields.ProjectId, query.ProjectId)
	}
	if len(query.ContractFullName) > 0 {
		mc.Like(dal.CustomerFields.ContactFullName, query.ContractFullName, daog.LikeStyleAll)
	}
	if len(query.OfficeCountry) > 0 {
		mc.Eq(dal.CustomerFields.OfficeCountry, query.OfficeCountry)
	}
	return mc
}

func (c *customer) Insert(ctx *tlog.Ctx, cus *dal.Customer) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	if err != nil {
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	exist, err := repo.Customer.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.CustomerFields.FullName, cus.FullName))
	if err != nil || exist != nil {
		return 0, errs.CompanyNameExist
	}
	//Insert customer
	_, err = repo.Customer.Insert(tc, cus)
	if err != nil {
		return 0, err
	}
	//Insert default group
	err = Group.createDefault(ctx, tc, domains.Business, cus.Id)
	if err != nil {
		return 0, err
	}
	return cus.Id, nil
}

func (c *customer) Update(ctx *tlog.Ctx, cus *dal.Customer) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	if err != nil {
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	id, err := repo.Customer.Update(tc, cus)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *customer) FindById(ctx *tlog.Ctx, id int64) (*dal.Customer, error) {
	ctx.Infof("customer find by id %d", id)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	if err != nil {
		ctx.Errorf("find customer new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	ret, err := repo.Customer.GetById(tc, id)
	if err != nil {
		ctx.Errorf("find customer by id error: %v", err)
		return nil, err
	}

	return ret, nil
}

func (c *customer) FindByIds(ctx *tlog.Ctx, ids []int64) ([]*dal.Customer, error) {
	ctx.Infof("customer find by ids %v", ids)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	if err != nil {
		ctx.Errorf("find customer new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	ret, err := repo.Customer.GetByIds(tc, ids)
	if err != nil {
		ctx.Errorf("find customer error: %v", err)
		return nil, err
	}
	return ret, nil
}

func (c *customer) FindAll(ctx *tlog.Ctx) ([]*dal.Customer, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			ctx.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()
	all, err := repo.Customer.GetAll(tc)
	if err != nil {
		return nil, err
	}
	return all, nil
}
