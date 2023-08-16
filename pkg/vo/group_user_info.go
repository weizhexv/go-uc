package vo

import "encoding/json"

type GroupUserInfo struct {
	Id       int64  `json:"id,omitempty"`
	Uid      int64  `json:"uid,omitempty"`
	GroupId  int64  `json:"groupId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
}

func (i *GroupUserInfo) String() string {
	if j, err := json.Marshal(i); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
