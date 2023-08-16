package vo

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
)

type UserInfo struct {
	Uid       int64                 `json:"uid"`
	Name      string                `json:"name"`
	Mobile    string                `json:"mobile,omitempty"`
	Email     string                `json:"email"`
	Nickname  string                `json:"nickname,omitempty"`
	Avatar    string                `json:"avatar,omitempty"`
	Domain    domains.Domain        `json:"domain,omitempty"`
	DomainId  int64                 `json:"domainId,omitempty"`
	Role      *roles.Role           `json:"role,omitempty"`
	GroupRole *roles.Role           `json:"groupRole,omitempty"`
	Enabled   bool                  `json:"enabled"`
	Deleted   bool                  `json:"deleted"`
	Verified  bool                  `json:"verified"`
	CreatedAt ttypes.NormalDatetime `json:"createdAt,omitempty"`
}

func (u *UserInfo) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
