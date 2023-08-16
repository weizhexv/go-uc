package model

type SupplierQuery struct {
	Id               int64  `json:"id,omitempty"`
	FullName         string `json:"fullName,omitempty"`
	ContractFullName string `json:"contractFullName,omitempty"`
	ServingCountries string `json:"servingCountries,omitempty"`
	PageNo           int    `json:"pageNo,omitempty"`
	PageSize         int    `json:"pageSize,omitempty"`
}
