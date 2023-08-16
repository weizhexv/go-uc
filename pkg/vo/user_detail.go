package vo

import (
	dgctx "dghire.com/libs/go-common/context"
	"encoding/json"
)

type UserDetail struct {
	*UserInfo
	DomainName string `json:"domainName"`
	Me         bool   `json:"me"`
}

func (u *UserDetail) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(j)
	}
}

func UserDetailOf(ctx *dgctx.DgContext, info *UserInfo, domainName string) *UserDetail {
	var myUid int64
	if ctx != nil {
		myUid = ctx.UserId
	}
	return &UserDetail{
		UserInfo:   info,
		DomainName: domainName,
		Me:         myUid == info.Uid,
	}
}

func UserDetailsOf(ctx *dgctx.DgContext, userInfos []*UserInfo, domainNames map[int64]string) []*UserDetail {
	var myUid int64
	if ctx != nil {
		myUid = ctx.UserId
	}
	var userDetails []*UserDetail
	if len(userInfos) == 0 {
		return userDetails
	}
	for _, e := range userInfos {
		userDetails = append(userDetails, &UserDetail{
			UserInfo:   e,
			DomainName: domainNames[e.DomainId],
			Me:         myUid == e.Uid,
		})
	}
	return userDetails
}
