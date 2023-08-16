package req

import (
	"encoding/json"
	"go-uc/internal/model"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/times"
	"go-uc/pkg/var/domains"
)

type SignUpParams struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Country    string `json:"country" form:"country"`
	TaxCountry string `json:"taxCountry" form:"taxCountry"`
	Birthday   string `json:"birthday" form:"birthday"`
	Passport   string `json:"passport" form:"passport"`
	IdCard     string `json:"idCard" form:"idCard"`
	Mobile     string `json:"mobile" form:"mobile"`
	Address    string `json:"address" form:"address"`
	PostalCode string `json:"postalCode" form:"postalCode"`
	Nonce      string `json:"nonce" form:"nonce" binding:"required"`
	CardType   string `json:"cardType" form:"cardType"`
}

func (p *SignUpParams) ToInsertUser(c *tlog.Ctx) *model.UpsertUser {
	uu := model.NewUpsertUser(domains.Employee, true)
	uu.Name = p.Name
	uu.Email = p.Email
	uu.Password = p.Password
	uu.Country = p.Country
	uu.TaxCountry = p.TaxCountry
	if len(p.Birthday) > 0 {
		if t, err := times.ParseNilableDate(p.Birthday); err != nil {
			c.Errorf("to insert user parse birthday err: %v", err)
		} else {
			uu.Birthday = t
		}
	}
	uu.Passport = p.Passport
	uu.IdCard = p.IdCard
	uu.Mobile = p.Mobile
	uu.Address = p.Address
	uu.PostalCode = p.PostalCode
	uu.CardType = p.CardType
	return uu
}

func (p *SignUpParams) String() string {
	if j, err := json.Marshal(p); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
