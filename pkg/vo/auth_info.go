package vo

import "encoding/json"

type AuthInfo struct {
	TokenInfo *TokenInfo `json:"tokenInfo,omitempty"`
	UserInfo  *UserInfo  `json:"userInfo,omitempty"`
}

func NewAuthInfo(token *TokenInfo, user *UserInfo) *AuthInfo {
	return &AuthInfo{
		TokenInfo: token,
		UserInfo:  user,
	}
}

func (i *AuthInfo) String() string {
	j, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
