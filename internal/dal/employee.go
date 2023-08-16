package dal

import (
	"encoding/json"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/vo"
)

var EmployeeFields = struct {
	Id           string
	Uid          string
	Name         string
	Mobile       string
	Email        string
	Password     string
	Nickname     string
	Avatar       string
	Country      string
	TaxCountry   string
	Birthday     string
	Passport     string
	IdCard       string
	Address      string
	PostalCode   string
	Verified     string
	BankAccounts string
	Attachments  string
	Seed         string
	CreatedBy    string
	CreatedAt    string
	ModifiedBy   string
	ModifiedAt   string
	Enabled      string
	Deleted      string
	AccountId    string
	CardType     string
}{
	"id",
	"uid",
	"name",
	"mobile",
	"email",
	"password",
	"nickname",
	"avatar",
	"country",
	"tax_country",
	"birthday",
	"passport",
	"id_card",
	"address",
	"postal_code",
	"verified",
	"bank_accounts",
	"attachments",
	"seed",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
	"enabled",
	"deleted",
	"account_id",
	"card_type",
}

var EmployeeMeta = &daog.TableMeta[Employee]{
	Table: "employee",
	Columns: []string{
		"id",
		"uid",
		"name",
		"mobile",
		"email",
		"password",
		"nickname",
		"avatar",
		"country",
		"tax_country",
		"birthday",
		"passport",
		"id_card",
		"address",
		"postal_code",
		"verified",
		"bank_accounts",
		"attachments",
		"seed",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
		"enabled",
		"deleted",
		"account_id",
		"card_type",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *Employee, point bool) any {
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
		if "country" == columnName {
			if point {
				return &ins.Country
			}
			return ins.Country
		}
		if "tax_country" == columnName {
			if point {
				return &ins.TaxCountry
			}
			return ins.TaxCountry
		}
		if "birthday" == columnName {
			if point {
				return &ins.Birthday
			}
			return ins.Birthday
		}
		if "passport" == columnName {
			if point {
				return &ins.Passport
			}
			return ins.Passport
		}
		if "id_card" == columnName {
			if point {
				return &ins.IdCard
			}
			return ins.IdCard
		}
		if "address" == columnName {
			if point {
				return &ins.Address
			}
			return ins.Address
		}
		if "postal_code" == columnName {
			if point {
				return &ins.PostalCode
			}
			return ins.PostalCode
		}
		if "verified" == columnName {
			if point {
				return &ins.Verified
			}
			return ins.Verified
		}
		if "bank_accounts" == columnName {
			if point {
				return &ins.BankAccounts
			}
			return ins.BankAccounts
		}
		if "attachments" == columnName {
			if point {
				return &ins.Attachments
			}
			return ins.Attachments
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
		if "account_id" == columnName {
			if point {
				return &ins.AccountId
			}
			return ins.AccountId
		}
		if "card_type" == columnName {
			if point {
				return &ins.CardType
			}
			return ins.CardType
		}
		return nil
	},
}

type Employee struct {
	Id           int64
	Uid          int64
	Name         string
	Mobile       ttypes.NilableString
	Email        string
	Password     ttypes.NilableString
	Nickname     ttypes.NilableString
	Avatar       ttypes.NilableString
	Country      ttypes.NilableString
	TaxCountry   ttypes.NilableString
	Birthday     ttypes.NilableDate
	Passport     ttypes.NilableString
	IdCard       ttypes.NilableString
	Address      ttypes.NilableString
	PostalCode   ttypes.NilableString
	BankAccounts ttypes.NilableString
	Attachments  ttypes.NilableString
	Seed         string
	CreatedBy    int64
	CreatedAt    ttypes.NormalDatetime
	ModifiedBy   int64
	ModifiedAt   ttypes.NormalDatetime
	Enabled      bool
	Deleted      bool
	Verified     bool
	AccountId    int64
	CardType     ttypes.NilableString
}

func (e Employee) String() string {
	bts, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(bts)
}

func (e Employee) ToInfo() *vo.UserInfo {
	return &vo.UserInfo{
		Uid:       e.Uid,
		Name:      e.Name,
		Mobile:    e.Mobile.String,
		Email:     e.Email,
		Nickname:  e.Nickname.String,
		Avatar:    e.Avatar.String,
		Domain:    domains.Employee,
		Enabled:   e.Enabled,
		Deleted:   e.Deleted,
		Verified:  e.Verified,
		CreatedAt: e.CreatedAt,
	}
}

func (e Employee) ToModel() *model.UserModel {
	return &model.UserModel{
		Domain:   domains.Employee,
		Id:       e.Id,
		Uid:      e.Uid,
		Password: e.Password.String,
		Seed:     []byte(e.Seed),
	}
}

func (e Employee) ToEmployeeInfo() (*vo.EmployeeInfo, error) {
	accounts := new([]*vo.BankAccount)
	if e.BankAccounts.Valid {
		err := json.Unmarshal([]byte(e.BankAccounts.String), accounts)
		if err != nil {
			return nil, err
		}
	}
	files := new([]*vo.FgwFile)
	if e.Attachments.Valid {
		err := json.Unmarshal([]byte(e.Attachments.String), files)
		if err != nil {
			return nil, err
		}
	}
	return &vo.EmployeeInfo{
		Uid:          e.Uid,
		Name:         e.Name,
		Mobile:       e.Mobile.String,
		Email:        e.Email,
		Domain:       domains.Employee.String(),
		Country:      e.Country.String,
		TaxCountry:   e.TaxCountry.String,
		Birthday:     e.Birthday,
		Passport:     e.Passport.String,
		IdCard:       e.IdCard.String,
		Address:      e.Address.String,
		PostalCode:   e.PostalCode.String,
		BankAccounts: accounts,
		Attachments:  files,
		Verified:     e.Verified,
		AccountId:    e.AccountId,
		CardType:     e.CardType.String,
	}, nil
}
