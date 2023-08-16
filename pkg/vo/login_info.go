package vo

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
)

type LoginInfo struct {
	Id         int64
	Uid        int64
	Domain     string
	Platform   string
	Identity   string
	ExpiredAt  ttypes.NormalDatetime
	CreatedBy  int64
	CreatedAt  ttypes.NormalDatetime
	ModifiedBy int64
	ModifiedAt ttypes.NormalDatetime
}

func (i *LoginInfo) String() string {
	if j, err := json.Marshal(i); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
