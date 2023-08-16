package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"time"
)

var UserIndexFields = struct {
	Id         string
	SubId      string
	Domain     string
	Identity   string
	CreatedBy  string
	CreatedAt  string
	ModifiedBy string
	ModifiedAt string
}{
	"id",
	"sub_id",
	"domain",
	"identity",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
}

var UserIndexMeta = &daog.TableMeta[UserIndex]{
	Table: "user_index",
	Columns: []string{
		"id",
		"sub_id",
		"domain",
		"identity",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *UserIndex, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "sub_id" == columnName {
			if point {
				return &ins.SubId
			}
			return ins.SubId
		}
		if "domain" == columnName {
			if point {
				return &ins.Domain
			}
			return ins.Domain
		}
		if "identity" == columnName {
			if point {
				return &ins.Identity
			}
			return ins.Identity
		}
		if "created_by" == columnName {
			if point {
				return &ins.CreatedBy
			}
			return ins.CreatedBy
		}
		if "created_at" == columnName {
			if point {
				return &ins.CreatedAt
			}
			return ins.CreatedAt
		}
		if "modified_by" == columnName {
			if point {
				return &ins.ModifiedBy
			}
			return ins.ModifiedBy
		}
		if "modified_at" == columnName {
			if point {
				return &ins.ModifiedAt
			}
			return ins.ModifiedAt
		}
		return nil
	},
}

type UserIndex struct {
	Id         int64
	SubId      int64
	Domain     string
	Identity   string
	CreatedBy  int64
	CreatedAt  ttypes.NilableDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NilableDatetime
}

func NewUserIndex(ctx *tlog.Ctx, domain domains.Domain, email string) *UserIndex {
	return &UserIndex{
		SubId:      0,
		Domain:     domain.String(),
		Identity:   email,
		CreatedBy:  ctx.UserId,
		CreatedAt:  *ttypes.FromDatetime(time.Now()),
		ModifiedBy: ctx.UserId,
		ModifiedAt: *ttypes.FromDatetime(time.Now()),
	}
}
