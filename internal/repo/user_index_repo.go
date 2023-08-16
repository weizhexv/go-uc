package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var UserIndex = userIndex{daog.NewBaseQuickDao(dal.UserIndexMeta)}

type userIndex struct {
	daog.QuickDao[dal.UserIndex]
}
