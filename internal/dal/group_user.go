package dal

import (
	"encoding/json"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/times"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

var GroupUserFields = struct {
	Id         string
	Uid        string
	GroupId    string
	RoleId     string
	CreatedBy  string
	CreatedAt  string
	ModifiedBy string
	ModifiedAt string
}{
	"id",
	"uid",
	"group_id",
	"role_id",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
}

var GroupUserMeta = &daog.TableMeta[GroupUser]{
	Table: "group_user",
	Columns: []string{
		"id",
		"uid",
		"group_id",
		"role_id",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *GroupUser, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "uid" == columnName {
			if point {
				return &ins.Uid
			}
			return ins.Uid
		}
		if "group_id" == columnName {
			if point {
				return &ins.GroupId
			}
			return ins.GroupId
		}
		if "role_id" == columnName {
			if point {
				return &ins.RoleId
			}
			return ins.RoleId
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

type GroupUser struct {
	Id         int64
	Uid        int64
	GroupId    int64
	RoleId     int
	CreatedBy  int64
	CreatedAt  ttypes.NormalDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NormalDatetime
}

func (gu *GroupUser) String() string {
	js, err := json.Marshal(gu)
	if err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (gu *GroupUser) ToInfo() *vo.GroupUserInfo {
	return &vo.GroupUserInfo{
		Id:       gu.Id,
		Uid:      gu.Uid,
		GroupId:  gu.GroupId,
		RoleName: roles.OfId(gu.RoleId).Name,
	}
}

func NewGroupUser(ctx *tlog.Ctx, upsert *model.UpsertUser) *GroupUser {
	return &GroupUser{
		Uid:        upsert.Uid,
		GroupId:    upsert.GroupId,
		RoleId:     upsert.GroupRole.Id,
		CreatedBy:  ctx.UserId,
		CreatedAt:  times.NowDatetime(),
		ModifiedBy: ctx.UserId,
		ModifiedAt: times.NowDatetime(),
	}
}
