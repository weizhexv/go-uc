package req

import (
	"encoding/json"
	"go-uc/internal/model"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/typeutil"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
)

type UpsertUserParams struct {
	Domain   string `json:"domain,omitempty" form:"domain" binding:"required"`
	DomainId int64  `json:"domainId,omitempty" form:"domainId"`
	Uid      int64  `json:"uid,omitempty" form:"uid"`
	Name     string `json:"name,omitempty" form:"name"`
	Mobile   string `json:"mobile,omitempty" form:"mobile"`
	Email    string `json:"email,omitempty" form:"email"`
	Role     string `json:"role,omitempty" form:"role"`
	Enabled  any    `json:"enabled,omitempty" form:"enabled"`
}

func (p *UpsertUserParams) ToUpsertUser(ctx *tlog.Ctx) (*model.UpsertUser, error) {
	enabled, checkEnabled := typeutil.GetBool(p.Enabled)
	var groupRole, role *roles.Role
	if len(p.Role) > 0 {
		if domains.Parse(p.Domain).Is(domains.Business) {
			if ctx.Domain().Is(domains.Platform) {
				role = roles.OfName(p.Role)
			} else if ctx.Domain().Is(domains.Business) {
				groupRole = roles.OfName(p.Role)
			}
		} else {
			role = roles.OfName(p.Role)
		}
	}

	domainId := p.DomainId
	if ctx.Domain().IsAny(domains.Business, domains.Supplier) {
		domainId = ctx.DomainId
	}

	return &model.UpsertUser{
		Insert:       p.Uid == 0,
		Uid:          p.Uid,
		Domain:       domains.Parse(p.Domain),
		DomainId:     domainId,
		Name:         p.Name,
		Mobile:       p.Mobile,
		Email:        p.Email,
		Role:         role,
		GroupRole:    groupRole,
		CheckEnabled: checkEnabled,
		Enabled:      enabled,
		GroupId:      ctx.GroupId,
	}, nil
}

func (p *UpsertUserParams) String() string {
	if j, err := json.Marshal(p); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
