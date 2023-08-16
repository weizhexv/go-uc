package model

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"go-uc/pkg/vo"
)

type UpsertUser struct {
	Insert bool  `json:"insert,omitempty"`
	Id     int64 `json:"id,omitempty"`
	Uid    int64 `json:"uid,omitempty"`

	Domain       domains.Domain `json:"domain,omitempty"`
	DomainId     int64          `json:"domainId,omitempty"`
	Name         string         `json:"name,omitempty"`
	Mobile       string         `json:"mobile,omitempty"`
	Email        string         `json:"email,omitempty"`
	Password     string         `json:"password,omitempty"`
	Nickname     string         `json:"nickname,omitempty"`
	Avatar       string         `json:"avatar,omitempty"`
	Role         *roles.Role    `json:"role,omitempty"`
	GroupRole    *roles.Role    `json:"groupRole,omitempty"`
	CheckEnabled bool           `json:"checkEnabled,omitempty"`
	Enabled      bool           `json:"enabled,omitempty"`
	GroupId      int64          `json:"groupId,omitempty"`

	//employee related
	Country       string              `json:"country,omitempty"`
	TaxCountry    string              `json:"taxCountry,omitempty"`
	Birthday      *ttypes.NilableDate `json:"birthday,omitempty"`
	Passport      string              `json:"passport,omitempty"`
	IdCard        string              `json:"idCard,omitempty"`
	Address       string              `json:"address,omitempty"`
	PostalCode    string              `json:"postalCode,omitempty"`
	CheckVerified bool                `json:"checkVerified,omitempty"`
	Verified      bool                `json:"verified,omitempty"`
	BankAccounts  []*vo.BankAccount   `json:"bankAccounts,omitempty"`
	Attachments   []*vo.FgwFile       `json:"attachments,omitempty"`
	AccountId     int64               `json:"accountId,omitempty"`
	CardType      string              `json:"cardType,omitempty"`
}

func NewUpsertUser(domain domains.Domain, insert bool) *UpsertUser {
	return &UpsertUser{Domain: domain, Insert: insert}
}

func (u *UpsertUser) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
