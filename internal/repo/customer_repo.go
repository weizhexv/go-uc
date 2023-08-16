package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Customer = customer{daog.NewBaseQuickDao(dal.CustomerMeta)}

type customer struct {
	daog.QuickDao[dal.Customer]
}
