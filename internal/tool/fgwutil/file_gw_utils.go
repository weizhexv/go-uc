package fgwutil

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"encoding/json"
	"go-uc/pkg/vo"
	"go-uc/vconfig"
	"strconv"
)

type FgwBizType string

const BizEmployee FgwBizType = "employee"

func AppendURL(files []*vo.FgwFile, bizType FgwBizType, bizId int64) (string, error) {
	if bizId <= 0 || len(bizType) == 0 {
		return "", dgerr.ARGUMENT_NOT_VALID
	}
	if len(files) == 0 {
		return "[]", nil
	}
	for _, file := range files {
		file.Url = NewURL(bizType, bizId, file.FileId)
	}
	if j, err := json.Marshal(files); err != nil {
		return "", err
	} else {
		return string(j), nil
	}
}

func NewURL(bizType FgwBizType, bizId int64, fileId int64) string {
	return vconfig.UrlFgw() + string(bizType) + "/" + strconv.FormatInt(bizId, 10) + "/?id=" + strconv.FormatInt(fileId, 10)
}
