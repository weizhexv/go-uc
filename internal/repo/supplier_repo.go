package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Supplier = supplier{daog.NewBaseQuickDao(dal.SupplierMeta)}

type supplier struct {
	daog.QuickDao[dal.Supplier]
}
