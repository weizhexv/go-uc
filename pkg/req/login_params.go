package req

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"encoding/json"
	"go-uc/internal/model"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/platforms"
)

type LoginParams struct {
	Platform string `json:"platform" form:"platform" binding:"required"`
	Domain   string `json:"domain" form:"domain" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (p *LoginParams) ToLoginContext() (*model.LoginContext, error) {
	plt := platforms.Parse(p.Platform)
	if len(plt) == 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}
	dm := domains.Parse(p.Domain)
	if len(dm) == 0 {
		return nil, dgerr.ARGUMENT_NOT_VALID
	}
	return &model.LoginContext{
		Platform: platforms.Parse(p.Platform),
		Domain:   dm,
		Email:    p.Email,
		Pwd:      p.Password,
	}, nil
}

func (p *LoginParams) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
