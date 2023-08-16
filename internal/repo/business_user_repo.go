package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
	"go-uc/internal/model"
)

var BusinessUser = businessUser{daog.NewBaseQuickDao(dal.BusinessUserMeta)}

type businessUser struct {
	daog.QuickDao[dal.BusinessUser]
}

func (b *businessUser) FindByUid(tc *daog.TransContext, uid int64) (*dal.BusinessUser, error) {
	mt := daog.NewMatcher().Eq(dal.BusinessUserFields.Uid, uid)
	ret, err := b.QueryOneMatcher(tc, mt)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (b *businessUser) FindCustomerSuperManagerCount(tc *daog.TransContext) ([]*model.IdCountPair, error) {
	f := func(ins *model.IdCountPair) []any {
		var sl []any
		sl = append(sl, ins.Id())
		sl = append(sl, ins.Count())
		return sl
	}
	sql := "select `customer_id` as `id`, count(id) from business_user where `enabled` = ? and `role_id` = ? group by `customer_id`"
	ret, err := daog.QueryRawSQL(tc, f, sql, 1, 4)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
