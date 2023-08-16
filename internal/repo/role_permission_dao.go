package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var RolePermission = rolePermission{daog.NewBaseQuickDao(dal.RolePermissionMeta)}

type rolePermission struct {
	daog.QuickDao[dal.RolePermission]
}
