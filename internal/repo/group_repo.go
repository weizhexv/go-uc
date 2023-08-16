package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Group = group{daog.NewBaseQuickDao(dal.GroupMeta)}

type group struct {
	daog.QuickDao[dal.Group]
}
