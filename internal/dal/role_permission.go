package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
)

var RolePermissionFields = struct {
	Id         string
	RoleId     string
	Domain     string
	Permission string
	CreatedAt  string
}{
	"id",
	"role_id",
	"domain",
	"permission",
	"created_at",
}

var RolePermissionMeta = &daog.TableMeta[RolePermission]{
	Table: "role_permission",
	Columns: []string{
		"id",
		"role_id",
		"domain",
		"permission",
		"created_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *RolePermission, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "role_id" == columnName {
			if point {
				return &ins.RoleId
			}
			return ins.RoleId
		}
		if "domain" == columnName {
			if point {
				return &ins.Domain
			}
			return ins.Domain
		}
		if "permission" == columnName {
			if point {
				return &ins.Permission
			}
			return ins.Permission
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

type RolePermission struct {
	Id         int64
	RoleId     int64
	Domain     string
	Permission string
	CreatedAt  ttypes.NormalDatetime
}
