package vo

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
)

type LoginDetail struct {
	UserInfo     *UserDetail           `json:"userInfo,omitempty"`
	AuthToken    string                `json:"authToken,omitempty"`
	ExpiredAt    ttypes.NormalDatetime `json:"expiredAt"`
	GroupUsers   []*GroupUserInfo      `json:"groupUsers,omitempty"`
	DomainEnable bool                  `json:"domainEnable"`
}

func (d *LoginDetail) String() string {
	if j, err := json.Marshal(d); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
