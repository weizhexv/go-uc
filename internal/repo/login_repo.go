package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Login = login{daog.NewBaseQuickDao(dal.UserLoginMeta)}

type login struct {
	daog.QuickDao[dal.UserLogin]
}
