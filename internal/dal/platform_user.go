package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

var PlatformUserFields = struct {
	Id         string
	Uid        string
	RoleId     string
	Name       string
	Mobile     string
	Email      string
	Password   string
	Nickname   string
	Avatar     string
	Seed       string
	CreatedBy  string
	CreatedAt  string
	ModifiedBy string
	ModifiedAt string
	Enabled    string
	Deleted    string
}{
	"id",
	"uid",
	"role_id",
	"name",
	"mobile",
	"email",
	"password",
	"nickname",
	"avatar",
	"seed",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
	"enabled",
	"deleted",
}

var PlatformUserMeta = &daog.TableMeta[PlatformUser]{
	Table: "platform_user",
	Columns: []string{
		"id",
		"uid",
		"role_id",
		"name",
		"mobile",
		"email",
		"password",
		"nickname",
		"avatar",
		"seed",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
		"enabled",
		"deleted",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *PlatformUser, point bool) any {
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
		if "nickname" == columnName {
			if point {
				return &ins.Nickname
			}
			return ins.Nickname
		}
		if "avatar" == columnName {
			if point {
				return &ins.Avatar
			}
			return ins.Avatar
		}
		if "seed" == columnName {
			if point {
				return &ins.Seed
			}
			return ins.Seed
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
		return nil
	},
}

type PlatformUser struct {
	Id         int64
	Uid        int64
	RoleId     int
	Name       string
	Mobile     ttypes.NilableString
	Email      string
	Password   ttypes.NilableString
	Nickname   ttypes.NilableString
	Avatar     ttypes.NilableString
	Seed       string
	CreatedBy  int64
	CreatedAt  ttypes.NormalDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NormalDatetime
	Enabled    bool
	Deleted    bool
}

func (u PlatformUser) ToInfo() *vo.UserInfo {
	return &vo.UserInfo{
		Uid:       u.Uid,
		Name:      u.Name,
		Mobile:    u.Mobile.StringNilAsEmpty(),
		Email:     u.Email,
		Nickname:  u.Nickname.StringNilAsEmpty(),
		Avatar:    u.Avatar.StringNilAsEmpty(),
		Domain:    domains.Platform,
		DomainId:  0,
		Role:      roles.OfId(u.RoleId),
		GroupRole: nil,
		Enabled:   u.Enabled,
		Deleted:   u.Deleted,
		CreatedAt: u.CreatedAt,
	}
}

func (u PlatformUser) ToModel() *model.UserModel {
	return &model.UserModel{
		Domain:   domains.Platform,
		Id:       u.Id,
		Uid:      u.Uid,
		Password: u.Password.String,
		Seed:     []byte(u.Seed),
	}
}
