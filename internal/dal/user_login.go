package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/pkg/var/platforms"
	"go-uc/pkg/vo"
	"time"
)

var UserLoginFields = struct {
	Id         string
	Uid        string
	Domain     string
	Platform   string
	Identity   string
	ExpiredAt  string
	CreatedBy  string
	CreatedAt  string
	ModifiedBy string
	ModifiedAt string
}{
	"id",
	"uid",
	"domain",
	"platform",
	"identity",
	"expired_at",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
}

var UserLoginMeta = &daog.TableMeta[UserLogin]{
	Table: "user_login",
	Columns: []string{
		"id",
		"uid",
		"domain",
		"platform",
		"identity",
		"expired_at",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *UserLogin, point bool) any {
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
		if "domain" == columnName {
			if point {
				return &ins.Domain
			}
			return ins.Domain
		}
		if "platform" == columnName {
			if point {
				return &ins.Platform
			}
			return ins.Platform
		}
		if "identity" == columnName {
			if point {
				return &ins.Identity
			}
			return ins.Identity
		}
		if "expired_at" == columnName {
			if point {
				return &ins.ExpiredAt
			}
			return ins.ExpiredAt
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

type UserLogin struct {
	Id         int64
	Uid        int64
	Domain     string
	Platform   string
	Identity   string
	ExpiredAt  ttypes.NormalDatetime
	CreatedBy  int64
	CreatedAt  ttypes.NormalDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NormalDatetime
}

func NewUserLogin(info *vo.UserInfo, plt platforms.Platform, period time.Duration) *UserLogin {
	return &UserLogin{
		Uid:        info.Uid,
		Domain:     info.Domain.String(),
		Platform:   plt.String(),
		Identity:   info.Email,
		ExpiredAt:  ttypes.NormalDatetime(time.Now().Add(period)),
		CreatedBy:  info.Uid,
		CreatedAt:  ttypes.NormalDatetime(time.Now()),
		ModifiedBy: info.Uid,
		ModifiedAt: ttypes.NormalDatetime(time.Now()),
	}
}

func (ul *UserLogin) ToInfo() *vo.LoginInfo {
	return &vo.LoginInfo{
		Id:         ul.Id,
		Uid:        ul.Uid,
		Domain:     ul.Domain,
		Platform:   ul.Platform,
		Identity:   ul.Identity,
		ExpiredAt:  ul.ExpiredAt,
		CreatedBy:  ul.CreatedBy,
		CreatedAt:  ul.CreatedAt,
		ModifiedBy: ul.ModifiedBy,
		ModifiedAt: ul.ModifiedAt,
	}
}
