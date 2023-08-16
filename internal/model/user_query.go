package model

import (
	"encoding/json"
	"github.com/rolandhe/daog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
)

type UserQuery struct {
	Domain       domains.Domain `json:"domain,omitempty"`
	DomainId     int64          `json:"domainId,omitempty"`
	Name         string         `json:"name,omitempty"`
	Mobile       string         `json:"mobile,omitempty"`
	Email        string         `json:"email,omitempty"`
	CheckEnabled bool           `json:"withEnabled,omitempty"`
	Enabled      bool           `json:"enabled,omitempty"`
	Roles        []*roles.Role  `json:"roles,omitempty"`
	GroupRole    *roles.Role    `json:"groupRole,omitempty"`
	GroupId      int64          `json:"groupId,omitempty"`
	PageNo       int            `json:"pageNo,omitempty"`
	PageSize     int            `json:"pageSize,omitempty"`
}

func (u *UserQuery) ToPager() *daog.Pager {
	return daog.NewPager(u.PageSize, u.PageNo)
}

func (u *UserQuery) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		return err.Error()
	}
	return string(j)
}
