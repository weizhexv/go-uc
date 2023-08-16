package req

import (
	"dghire.com/libs/go-common/page"
	"encoding/json"
	"go-uc/internal/model"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/typeutil"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
)

type QueryUserParams struct {
	Domain    string `json:"domain,omitempty" form:"domain" binding:"required"`
	DomainId  int64  `json:"domainId,omitempty" form:"domainId"`
	Name      string `json:"name,omitempty" form:"name"`
	Mobile    string `json:"mobile,omitempty" form:"mobile"`
	Email     string `json:"email,omitempty" form:"email"`
	Enabled   string `json:"enabled,omitempty" form:"enabled"`
	Role      string `json:"role,omitempty" form:"role"`
	GroupId   int64  `json:"groupId,omitempty" form:"groupId"`
	GroupRole string `json:"groupRole,omitempty" form:"groupRole"`
	page.PageParam
}

func (p *QueryUserParams) ToUserQuery(ctx *tlog.Ctx) *model.UserQuery {
	ctx.Infof("to user query group id %d", ctx.GroupId)
	enabled, checkEnabled := typeutil.GetBool(p.Enabled)
	var rs []*roles.Role
	if len(p.Role) > 0 {
		if roles.OfName(p.Role) != nil {
			rs = append(rs, roles.OfName(p.Role))
		}
	}
	groupId := p.GroupId
	if ctx.Domain().Is(domains.Business) && groupId == 0 {
		groupId = ctx.GroupId
	}

	return &model.UserQuery{
		Domain:       domains.Parse(p.Domain),
		DomainId:     p.DomainId,
		Name:         p.Name,
		Mobile:       p.Mobile,
		Email:        p.Email,
		CheckEnabled: checkEnabled,
		Enabled:      enabled,
		Roles:        rs,
		GroupRole:    roles.OfName(p.GroupRole),
		GroupId:      groupId,
		PageNo:       p.PageNo,
		PageSize:     p.PageSize,
	}
}

func (p *QueryUserParams) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
