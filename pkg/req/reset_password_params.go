package req

import "encoding/json"

type ResetPasswordParams struct {
	Email       string `json:"email,omitempty" form:"email" binding:"required"`
	Password    string `json:"password,omitempty" form:"password" binding:"required"`
	OldPassword string `json:"oldPassword,omitempty" form:"oldPassword"`
}

func (p *ResetPasswordParams) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
