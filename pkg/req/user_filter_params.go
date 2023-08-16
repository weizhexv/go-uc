package req

import "encoding/json"

type UserFilterParams struct {
	Domain    string `json:"domain,omitempty" form:"domain" binding:"required"`
	DomainId  int64  `json:"domainId,omitempty" form:"domainId"`
	Role      string `json:"role,omitempty" form:"role"`
	GroupId   int64  `json:"groupId,omitempty" form:"groupId"`
	GroupRole string `json:"groupRole,omitempty" form:"groupRole"`
}

func (p *UserFilterParams) String() string {
	if j, err := json.Marshal(p); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
