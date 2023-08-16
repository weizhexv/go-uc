package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var GroupUser = groupUser{daog.NewBaseQuickDao(dal.GroupUserMeta)}

type groupUser struct {
	daog.QuickDao[dal.GroupUser]
}

func (g *groupUser) FindByGroupIdAndUid(tc *daog.TransContext, groupId int64, uid int64) (*dal.GroupUser, error) {
	mc := daog.NewMatcher().Eq(dal.GroupUserFields.GroupId, groupId).Eq(dal.GroupUserFields.Uid, uid)
	ret, err := g.QueryOneMatcher(tc, mc)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
