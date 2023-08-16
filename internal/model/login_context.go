package model

import (
	"encoding/json"
	"go-uc/internal/tool/tokens"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/platforms"
	"go-uc/pkg/vo"
	"time"
)

type LoginContext struct {
	Platform platforms.Platform `json:"platform,omitempty"`
	Domain   domains.Domain     `json:"domain,omitempty"`
	Email    string             `json:"email,omitempty"`
	Pwd      string             `json:"pwd,omitempty"`

	Success bool  `json:"success"`
	Err     error `json:"err,omitempty"`

	UserDetail     *vo.UserDetail      `json:"userDetail,omitempty"`
	LoginInfo      *vo.LoginInfo       `json:"loginInfo,omitempty"`
	GroupUserInfos []*vo.GroupUserInfo `json:"groupUserInfos,omitempty"`
	Seed           []byte              `json:"seed,omitempty"`
	DomainEnable   bool                `json:"domainEnable"`
}

func (l *LoginContext) String() string {
	js, err := json.Marshal(l)
	if err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (l *LoginContext) Abort(err error) {
	l.Success = false
	l.Err = err
}

func (l *LoginContext) Done(info *vo.UserInfo, login *vo.LoginInfo, dmInfo *DomainInfo, GroupUserInfos []*vo.GroupUserInfo, seed []byte) {
	l.Success = true
	l.Err = nil
	l.UserDetail = vo.UserDetailOf(nil, info, dmInfo.DomainName)
	l.LoginInfo = login
	l.GroupUserInfos = GroupUserInfos
	l.Seed = seed
	l.DomainEnable = dmInfo.Enabled
}

func (l *LoginContext) ToLoginDetail() (*vo.LoginDetail, error) {
	tokenInfo := vo.NewTokenInfo(l.UserDetail.Uid, l.Domain.String(), l.Platform.String(), time.Time(l.LoginInfo.ExpiredAt))
	authToken, err := tokens.Sign(tokenInfo, l.Seed)
	if err != nil {
		return nil, err
	}
	return &vo.LoginDetail{
		UserInfo:     l.UserDetail,
		AuthToken:    authToken,
		ExpiredAt:    l.LoginInfo.ExpiredAt,
		GroupUsers:   l.GroupUserInfos,
		DomainEnable: l.DomainEnable,
	}, nil
}
