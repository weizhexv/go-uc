package req

import (
	"encoding/json"
	"fmt"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
	"go-uc/internal/tool/times"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/vo"
)

type UpdateEmployeeParams struct {
	Name         string            `form:"name" json:"name,omitempty"`
	Email        string            `form:"email" json:"email,omitempty"`
	Country      string            `form:"country" json:"country,omitempty"`
	TaxCountry   string            `form:"taxCountry" json:"taxCountry,omitempty"`
	Birthday     string            `form:"birthday" json:"birthday,omitempty"`
	IdCard       string            `form:"idCard" json:"idCard,omitempty"`
	Passport     string            `form:"passport" json:"passport,omitempty"`
	Mobile       string            `form:"mobile" json:"mobile,omitempty"`
	Address      string            `form:"address" json:"address,omitempty"`
	PostalCode   string            `form:"postalCode" json:"postalCode,omitempty"`
	BankAccounts []*vo.BankAccount `form:"bankAccounts" json:"bankAccounts,omitempty"`
	Attachments  []*vo.FgwFile     `form:"attachments" json:"attachments,omitempty"`
	AccountId    int64             `form:"accountId" json:"accountId,omitempty"`
	CardType     string            `form:"cardType" json:"cardType,omitempty"`
}

func (p *UpdateEmployeeParams) Validate() bool {
	if len(p.BankAccounts) > 0 {
		for _, account := range p.BankAccounts {
			if len(account.AccountNo) == 0 {
				return false
			}
			if len(account.OwnerName) == 0 {
				return false
			}
			if len(account.BankName) == 0 {
				return false
			}
		}
	}
	return true
}

func (p *UpdateEmployeeParams) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(j)
}

func (p *UpdateEmployeeParams) ToUpsertUser(uid int64) *model.UpsertUser {
	var birthday *ttypes.NilableDate
	if birth, err := times.ParseNilableDate(p.Birthday); err != nil {
		fmt.Println("parse nilable date error: ", err)
	} else {
		birthday = birth
	}
	return &model.UpsertUser{
		Insert:       false,
		Uid:          uid,
		Domain:       domains.Employee,
		Name:         p.Name,
		Mobile:       p.Mobile,
		Email:        p.Email,
		Country:      p.Country,
		TaxCountry:   p.TaxCountry,
		Birthday:     birthday,
		Passport:     p.Passport,
		IdCard:       p.IdCard,
		Address:      p.Address,
		PostalCode:   p.PostalCode,
		BankAccounts: p.BankAccounts,
		Attachments:  p.Attachments,
		AccountId:    p.AccountId,
		CardType:     p.CardType,
	}
}
