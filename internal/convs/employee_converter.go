package convs

import (
	dgctx "dghire.com/libs/go-common/context"
	dgerr "dghire.com/libs/go-common/enums/error"
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/tool/fgwutil"
	"go-uc/internal/tool/passwords"
	"go-uc/internal/tool/times"
)

type EmployeeConv struct {
	uu *model.UpsertUser
}

func NewEmployeeConv(uu *model.UpsertUser) *EmployeeConv {
	return &EmployeeConv{uu}
}

func (c *EmployeeConv) ToPO(ctx *dgctx.DgContext) (*dal.Employee, error) {
	hash, err := passwords.Hash(c.uu.Password)
	if err != nil {
		return nil, err
	}

	birthday := ttypes.NilableDate{}
	if c.uu.Birthday != nil {
		birthday = *c.uu.Birthday
	}
	return &dal.Employee{
		Uid:          c.uu.Uid,
		Name:         c.uu.Name,
		Mobile:       *ttypes.FromString(c.uu.Mobile),
		Email:        c.uu.Email,
		Password:     *ttypes.FromString(hash),
		Nickname:     *ttypes.FromString(c.uu.Nickname),
		Avatar:       *ttypes.FromString(c.uu.Avatar),
		Country:      *ttypes.FromString(c.uu.Country),
		TaxCountry:   *ttypes.FromString(c.uu.TaxCountry),
		Birthday:     birthday,
		Passport:     *ttypes.FromString(c.uu.Passport),
		IdCard:       *ttypes.FromString(c.uu.IdCard),
		Address:      *ttypes.FromString(c.uu.Address),
		PostalCode:   *ttypes.FromString(c.uu.PostalCode),
		BankAccounts: *ttypes.FromString("[]"),
		Attachments:  *ttypes.FromString("[]"),
		Seed:         passwords.RandomSeed(),
		CreatedBy:    ctx.UserId,
		CreatedAt:    times.NowDatetime(),
		ModifiedBy:   ctx.UserId,
		ModifiedAt:   times.NowDatetime(),
		Enabled:      true,
		Deleted:      false,
		Verified:     false,
		CardType:     *ttypes.FromString(c.uu.CardType),
	}, nil
}

func (c *EmployeeConv) UpdatePO(ctx *dgctx.DgContext, po *dal.Employee) error {
	if po == nil {
		return dgerr.ILLEGAL_OPERATION
	}
	if len(c.uu.Name) > 0 {
		po.Name = c.uu.Name
	}
	if len(c.uu.Mobile) > 0 {
		po.Mobile = *ttypes.FromString(c.uu.Mobile)
	}
	if len(c.uu.Password) > 0 {
		pwd, err := passwords.Hash(c.uu.Password)
		if err != nil {
			return err
		}
		po.Password = *ttypes.FromString(pwd)
	}
	if len(c.uu.Nickname) > 0 {
		po.Nickname = *ttypes.FromString(c.uu.Nickname)
	}
	if len(c.uu.Avatar) > 0 {
		po.Avatar = *ttypes.FromString(c.uu.Avatar)
	}
	if len(c.uu.Country) > 0 {
		po.Country = *ttypes.FromString(c.uu.Country)
	}
	if len(c.uu.TaxCountry) > 0 {
		po.TaxCountry = *ttypes.FromString(c.uu.TaxCountry)
	}
	if c.uu.Birthday != nil {
		po.Birthday = *c.uu.Birthday
	}
	if len(c.uu.Passport) > 0 {
		po.Passport = *ttypes.FromString(c.uu.Passport)
	}
	if len(c.uu.IdCard) > 0 {
		po.IdCard = *ttypes.FromString(c.uu.IdCard)
	}
	if len(c.uu.Address) > 0 {
		po.Address = *ttypes.FromString(c.uu.Address)
	}
	if len(c.uu.PostalCode) > 0 {
		po.PostalCode = *ttypes.FromString(c.uu.PostalCode)
	}
	if c.uu.CheckVerified {
		po.Verified = c.uu.Verified
	}
	if len(c.uu.BankAccounts) > 0 {
		j, err := json.Marshal(c.uu.BankAccounts)
		if err != nil {
			return err
		}
		po.BankAccounts = *ttypes.FromString(string(j))
	}
	if len(c.uu.Attachments) > 0 {
		files, err := fgwutil.AppendURL(c.uu.Attachments, fgwutil.BizEmployee, po.Id)
		if err != nil {
			return err
		}
		po.Attachments = *ttypes.FromString(files)
	}
	po.ModifiedAt = times.NowDatetime()
	po.ModifiedBy = ctx.UserId
	if c.uu.CheckEnabled {
		po.Enabled = c.uu.Enabled
	}
	if c.uu.AccountId > 0 {
		po.AccountId = c.uu.AccountId
	}
	if len(c.uu.CardType) > 0 {
		po.CardType = *ttypes.FromString(c.uu.CardType)
	}
	return nil
}
