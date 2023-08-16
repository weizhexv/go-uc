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
	"go-uc/pkg/var/errs"
)

type supplier struct {
}

var Supplier = new(supplier)

func (s *supplier) Query(c *tlog.Ctx, query *model.SupplierQuery) ([]*dal.Supplier, error) {
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
	mc := s.queryToMatcher(query)
	var pg *daog.Pager
	if query.PageNo > 0 && query.PageSize > 0 {
		pg = daog.NewPager(query.PageNo, query.PageSize)
	}
	ret, err := repo.Supplier.QueryPageListMatcher(tc, mc, pg)
	return ret, err
}

func (s *supplier) Count(c *tlog.Ctx, query *model.SupplierQuery) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
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
	ct, err := repo.Supplier.Count(tc, s.queryToMatcher(query))
	return ct, err
}

func (s *supplier) Insert(c *tlog.Ctx, spr *dal.Supplier) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
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
	exist, err := repo.Supplier.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.SupplierFields.FullName, spr.FullName))
	if err != nil {
		return 0, err
	}
	if exist != nil {
		return 0, errs.SupplierExist
	}
	_, err = repo.Supplier.Insert(tc, spr)
	return spr.Id, err
}

func (s *supplier) Update(c *tlog.Ctx, spr *dal.Supplier) (int64, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, c.TraceId)
	if err != nil {
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
	id, err := repo.Supplier.Update(tc, spr)
	return id, err
}

func (s *supplier) FindById(c *tlog.Ctx, id int64) (*dal.Supplier, error) {
	c.Infof("supplier find by id %d", id)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find supplier new trans error: %v", err)
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
	ret, err := repo.Supplier.GetById(tc, id)
	if err != nil {
		c.Errorf("find supplier error: %v", err)
	}
	return ret, err
}

func (s *supplier) FindAll(c *tlog.Ctx) ([]*dal.Supplier, error) {
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
	all, err := repo.Supplier.GetAll(tc)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (s *supplier) FindByIds(c *tlog.Ctx, ids []int64) ([]*dal.Supplier, error) {
	c.Infof("find supplier by ids %v", ids)
	if len(ids) == 0 {
		return []*dal.Supplier{}, nil
	}
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("find supplier new trans error: %v", err)
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
	ret, err := repo.Supplier.GetByIds(tc, ids)
	if err != nil {
		c.Errorf("find supplier by ids error: %v", err)
		return nil, err
	}
	return ret, err
}

func (s *supplier) queryToMatcher(query *model.SupplierQuery) daog.Matcher {
	mc := daog.NewMatcher()
	if query.Id != 0 {
		mc.Eq(dal.SupplierFields.Id, query.Id)
	}
	if len(query.FullName) > 0 {
		mc.Like(dal.SupplierFields.FullName, query.FullName, daog.LikeStyleAll)
	}
	if len(query.ContractFullName) > 0 {
		mc.Like(dal.SupplierFields.ContactFullName, query.ContractFullName, daog.LikeStyleAll)
	}
	if len(query.ServingCountries) > 0 {
		mc.Like(dal.SupplierFields.ServingCountries, `"`+query.ServingCountries+`"`, daog.LikeStyleAll)
	}
	return mc
}
