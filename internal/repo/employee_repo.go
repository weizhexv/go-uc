package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Employee = employee{daog.NewBaseQuickDao(dal.EmployeeMeta)}

type employee struct {
	daog.QuickDao[dal.Employee]
}

func (e *employee) FindByUid(tc *daog.TransContext, uid int64) (*dal.Employee, error) {
	mc := daog.NewMatcher().Eq(dal.EmployeeFields.Uid, uid)
	ret, err := e.QueryOneMatcher(tc, mc)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
