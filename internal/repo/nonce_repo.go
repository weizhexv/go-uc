package repo

import (
	"github.com/rolandhe/daog"
	"go-uc/internal/dal"
)

var Nonce = nonce{daog.NewBaseQuickDao(dal.NonceMeta)}

type nonce struct {
	daog.QuickDao[dal.Nonce]
}
