package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

var BusinessUserFields = struct {
	Id         string
	Uid        string
	CustomerId string
	RoleId     string
	Name       string
	Avatar     string
	Nickname   string
	Mobile     string
	Email      string
	Password   string
	CreatedBy  string
	CreatedAt  string
	ModifiedBy string
	ModifiedAt string
	Enabled    string
	Deleted    string
	Vsn        string
	Seed       string
}{
	"id",
	"uid",
	"customer_id",
	"role_id",
	"name",
	"avatar",
	"nickname",
	"mobile",
	"email",
	"password",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
	"enabled",
	"deleted",
	"vsn",
	"seed",
}

var BusinessUserMeta = &daog.TableMeta[BusinessUser]{
	Table: "business_user",
	Columns: []string{
		"id",
		"uid",
		"customer_id",
		"role_id",
		"name",
		"avatar",
		"nickname",
		"mobile",
		"email",
		"password",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
		"enabled",
		"deleted",
		"vsn",
		"seed",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *BusinessUser, point bool) any {
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
		if "customer_id" == columnName {
			if point {
				return &ins.CustomerId
			}
			return ins.CustomerId
		}
		if "role_id" == columnName {
			if point {
				return &ins.RoleId
			}
			return ins.RoleId
		}
		if "name" == columnName {
			if point {
				return &ins.Name
			}
			return ins.Name
		}
		if "avatar" == columnName {
			if point {
				return &ins.Avatar
			}
			return ins.Avatar
		}
		if "nickname" == columnName {
			if point {
				return &ins.Nickname
			}
			return ins.Nickname
		}
		if "mobile" == columnName {
			if point {
				return &ins.Mobile
			}
			return ins.Mobile
		}
		if "email" == columnName {
			if point {
				return &ins.Email
			}
			return ins.Email
		}
		if "password" == columnName {
			if point {
				return &ins.Password
			}
			return ins.Password
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
		if "enabled" == columnName {
			if point {
				return &ins.Enabled
			}
			return ins.Enabled
		}
		if "deleted" == columnName {
			if point {
				return &ins.Deleted
			}
			return ins.Deleted
		}
		if "vsn" == columnName {
			if point {
				return &ins.Vsn
			}
			return ins.Vsn
		}
		if "seed" == columnName {
			if point {
				return &ins.Seed
			}
			return ins.Seed
		}

		return nil
	},
}

type BusinessUser struct {
	Id         int64
	Uid        int64
	CustomerId int64
	RoleId     int
	Name       string
	Avatar     ttypes.NilableString
	Nickname   ttypes.NilableString
	Mobile     ttypes.NilableString
	Email      string
	Password   ttypes.NilableString
	CreatedBy  int64
	CreatedAt  ttypes.NormalDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NormalDatetime
	Seed       string
	Enabled    bool
	Deleted    bool
	Vsn        int64
}

func (b *BusinessUser) ToInfo(gRoles map[int64]*roles.Role) *vo.UserInfo {
	var gRole *roles.Role
	if len(gRoles) > 0 {
		gRole = gRoles[b.Uid]
	}
	return &vo.UserInfo{
		Uid:       b.Uid,
		Name:      b.Name,
		Mobile:    b.Mobile.String,
		Email:     b.Email,
		Nickname:  b.Nickname.String,
		Avatar:    b.Avatar.String,
		Domain:    domains.Business,
		DomainId:  b.CustomerId,
		Role:      roles.OfId(b.RoleId),
		GroupRole: gRole,
		Enabled:   b.Enabled,
		Deleted:   b.Deleted,
		Verified:  false,
		CreatedAt: b.CreatedAt,
	}
}

func (b *BusinessUser) ToModel() *model.UserModel {
	return &model.UserModel{
		Domain:   domains.Business,
		Id:       b.Id,
		Uid:      b.Uid,
		Password: b.Password.String,
		Seed:     []byte(b.Seed),
	}
}
