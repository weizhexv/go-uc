package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/tool/stringutil"
	"go-uc/internal/tool/times"
	"time"
)

var NonceFields = struct {
	Id        string
	Identity  string
	Nonce     string
	Status    string
	CreatedBy string
	ExpireAt  string
	CreatedAt string
}{
	"id",
	"identity",
	"nonce",
	"status",
	"created_by",
	"expire_at",
	"created_at",
}

var NonceMeta = &daog.TableMeta[Nonce]{
	Table: "nonce",
	Columns: []string{
		"id",
		"identity",
		"nonce",
		"status",
		"created_by",
		"expire_at",
		"created_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *Nonce, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "identity" == columnName {
			if point {
				return &ins.Identity
			}
			return ins.Identity
		}
		if "nonce" == columnName {
			if point {
				return &ins.Nonce
			}
			return ins.Nonce
		}
		if "status" == columnName {
			if point {
				return &ins.Status
			}
			return ins.Status
		}
		if "created_by" == columnName {
			if point {
				return &ins.CreatedBy
			}
			return ins.CreatedBy
		}
		if "expire_at" == columnName {
			if point {
				return &ins.ExpireAt
			}
			return ins.ExpireAt
		}
		if "created_at" == columnName {
			if point {
				return &ins.CreatedAt
			}
			return ins.CreatedAt
		}

		return nil
	},
}

type Nonce struct {
	Id        int64
	Identity  string
	Nonce     string
	Status    int8
	CreatedBy int64
	ExpireAt  ttypes.NormalDatetime
	CreatedAt ttypes.NormalDatetime
}

func NewNonce(uid int64, identity string) *Nonce {
	return &Nonce{
		Identity:  identity,
		Nonce:     stringutil.Random(20),
		Status:    0,
		CreatedBy: uid,
		ExpireAt:  ttypes.NormalDatetime(time.Now().AddDate(0, 0, 30)),
		CreatedAt: times.NowDatetime(),
	}
}
