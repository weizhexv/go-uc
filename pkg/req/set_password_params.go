package req

import "encoding/json"

type SetPasswordParams struct {
	Email    string `json:"email,omitempty" form:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
	Nonce    string `json:"nonce,omitempty" form:"nonce" binding:"required"`
}

func (p *SetPasswordParams) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
