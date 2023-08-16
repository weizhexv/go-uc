package model

import "encoding/json"

type DomainInfo struct {
	DomainId   int64  `json:"domainId"`
	DomainName string `json:"domainName"`
	Enabled    bool   `json:"enabled"`
}

func (info *DomainInfo) String() string {
	jsn, err := json.Marshal(info)
	if err != nil {
		return err.Error()
	}
	return string(jsn)
}
