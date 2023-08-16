package vo

import (
	"encoding/json"
	"github.com/rolandhe/daog/ttypes"
)

type EmployeeInfo struct {
	Uid                  int64                   `json:"uid,omitempty"`
	Name                 string                  `json:"name,omitempty"`
	Mobile               string                  `json:"mobile,omitempty"`
	Email                string                  `json:"email,omitempty"`
	Domain               string                  `json:"domain,omitempty"`
	Country              string                  `json:"country,omitempty"`
	TaxCountry           string                  `json:"taxCountry,omitempty"`
	Birthday             ttypes.NilableDate      `json:"birthday,omitempty"`
	Passport             string                  `json:"passport,omitempty"`
	IdCard               string                  `json:"idCard,omitempty"`
	Address              string                  `json:"address,omitempty"`
	PostalCode           string                  `json:"postalCode,omitempty"`
	BankAccounts         *[]*BankAccount         `json:"bankAccounts,omitempty"`
	Attachments          *[]*FgwFile             `json:"attachments,omitempty"`
	Verified             bool                    `json:"verified"`
	AccountId            int64                   `json:"accountId"`
	CardType             string                  `json:"cardType"`
	VerificationDecision *VerificationDecisionVo `json:"verificationDecision"`
}

type BankAccount struct {
	BankName  string `json:"bankName,omitempty"`
	AccountNo string `json:"accountNo,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
}

type FgwFile struct {
	FileId   int64  `json:"fileId,omitempty"`
	Filename string `json:"filename,omitempty"`
	Format   string `json:"format,omitempty"`
	Size     int64  `json:"size,omitempty"`
	Url      string `json:"url,omitempty"`
}

type VerificationDecisionVo struct {
	Status     string `json:"status,omitempty"`
	StatusName string `json:"statusName,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

func (i *EmployeeInfo) String() string {
	if j, err := json.Marshal(i); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
