package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var SupplierUser = supplierUser{daog.NewBaseQuickDao(dal.SupplierUserMeta)}

type supplierUser struct {
	daog.QuickDao[dal.SupplierUser]
}
