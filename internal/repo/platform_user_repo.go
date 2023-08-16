package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var PlatformUser = platformUser{daog.NewBaseQuickDao(dal.PlatformUserMeta)}

type platformUser struct {
	daog.QuickDao[dal.PlatformUser]
}
