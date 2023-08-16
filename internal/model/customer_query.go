package model

import "encoding/json"

type CustomerQuery struct {
	Id               int64  `json:"id,omitempty"`
	FullName         string `json:"fullName,omitempty"`
	CheckEnabled     bool   `json:"checkEnabled,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	ProjectId        string `json:"projectId,omitempty"`
	ContractFullName string `json:"contractFullName,omitempty"`
	OfficeCountry    string `json:"officeCountry,omitempty"`
	PageNo           int    `json:"pageNo,omitempty"`
	PageSize         int    `json:"pageSize,omitempty"`
}

func (q *CustomerQuery) String() string {
	bs, err := json.Marshal(q)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}
