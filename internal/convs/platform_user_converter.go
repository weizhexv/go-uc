package convs

import (
	dgctx "dghire.com/libs/go-common/context"
	dgerr "dghire.com/libs/go-common/enums/error"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/tool/passwords"
	"go-uc/internal/tool/times"
)

type PlatformUserConv struct {
	uu *model.UpsertUser
}

func NewPlatformUserConv(uu *model.UpsertUser) *PlatformUserConv {
	return &PlatformUserConv{uu}
}

func (c *PlatformUserConv) ToPO(ctx *dgctx.DgContext) *dal.PlatformUser {
	return &dal.PlatformUser{
		Uid:        c.uu.Uid,
		RoleId:     c.uu.Role.Id,
		Name:       c.uu.Name,
		Mobile:     *ttypes.FromString(c.uu.Mobile),
		Email:      c.uu.Email,
		Password:   ttypes.NilableString{},
		Nickname:   *ttypes.FromString(c.uu.Nickname),
		Avatar:     *ttypes.FromString(c.uu.Avatar),
		Seed:       passwords.RandomSeed(),
		CreatedBy:  ctx.UserId,
		CreatedAt:  times.NowDatetime(),
		ModifiedBy: ctx.UserId,
		ModifiedAt: times.NowDatetime(),
		Enabled:    true,
	}
}

func (c *PlatformUserConv) UpdatePO(ctx *dgctx.DgContext, po *dal.PlatformUser) error {
	if po == nil {
		return dgerr.ILLEGAL_OPERATION
	}
	if c.uu.Role != nil {
		po.RoleId = c.uu.Role.Id
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
	po.ModifiedAt = times.NowDatetime()
	po.ModifiedBy = ctx.UserId
	if c.uu.CheckEnabled {
		po.Enabled = c.uu.Enabled
	}
	return nil
}
