package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/internal/model"
)

var CustomerFields = struct {
	Id                       string
	ContractId               string
	FullName                 string
	ProjectId                string
	LegalStatus              string
	CustomerStatus           string
	TaxId                    string
	FieldOfBusiness          string
	OfficeCountry            string
	OfficeCity               string
	OfficeStreet             string
	OfficeRegion             string
	OfficePostalCode         string
	OfficeNo                 string
	AddressCountry           string
	AddressCity              string
	AddressStreet            string
	AddressTelephone         string
	AddressRegion            string
	AddressPostalCode        string
	AddressNo                string
	AddressFax               string
	ContactFullName          string
	ContactSurname           string
	ContactTelephone         string
	ContactFirstname         string
	ContactEmail             string
	ContactPosition          string
	BankName                 string
	BankAccountName          string
	Bankgiro                 string
	Iban                     string
	BankAddress              string
	BankAccountNo            string
	BicSwift                 string
	BankCurrency             string
	InvoiceConfirmFirstname  string
	InvoiceConfirmEmail      string
	InvoiceConfirmDuoDate    string
	InvoiceConfirmSurname    string
	InvoiceConfirmTelephone  string
	InvoiceConfirmCurrency   string
	InvoiceReceiverFirstname string
	InvoiceReceiverSurname   string
	InvoiceReceiverEmail     string
	InvoiceReceiverTelephone string
	Enabled                  string
	CreatedAt                string
	ModifiedAt               string
	CreatedBy                string
	ModifiedBy               string
}{
	"id",
	"contract_id",
	"full_name",
	"project_id",
	"legal_status",
	"customer_status",
	"tax_id",
	"field_of_business",
	"office_country",
	"office_city",
	"office_street",
	"office_region",
	"office_postal_code",
	"office_no",
	"address_country",
	"address_city",
	"address_street",
	"address_telephone",
	"address_region",
	"address_postal_code",
	"address_no",
	"address_fax",
	"contact_full_Name",
	"contact_surname",
	"contact_telephone",
	"contact_firstname",
	"contact_email",
	"contact_position",
	"bank_name",
	"bank_account_name",
	"bankgiro",
	"iban",
	"bank_address",
	"bank_account_no",
	"bic_swift",
	"bank_currency",
	"invoice_confirm_firstname",
	"invoice_confirm_email",
	"invoice_confirm_duo_date",
	"invoice_confirm_surname",
	"invoice_confirm_telephone",
	"invoice_confirm_currency",
	"invoice_receiver_firstname",
	"invoice_receiver_surname",
	"invoice_receiver_email",
	"invoice_receiver_telephone",
	"enabled",
	"created_at",
	"modified_at",
	"created_by",
	"modified_by",
}

var CustomerMeta = &daog.TableMeta[Customer]{
	Table: "customer",
	Columns: []string{
		"id",
		"contract_id",
		"full_name",
		"project_id",
		"legal_status",
		"customer_status",
		"tax_id",
		"field_of_business",
		"office_country",
		"office_city",
		"office_street",
		"office_region",
		"office_postal_code",
		"office_no",
		"address_country",
		"address_city",
		"address_street",
		"address_telephone",
		"address_region",
		"address_postal_code",
		"address_no",
		"address_fax",
		"contact_full_Name",
		"contact_surname",
		"contact_telephone",
		"contact_firstname",
		"contact_email",
		"contact_position",
		"bank_name",
		"bank_account_name",
		"bankgiro",
		"iban",
		"bank_address",
		"bank_account_no",
		"bic_swift",
		"bank_currency",
		"invoice_confirm_firstname",
		"invoice_confirm_email",
		"invoice_confirm_duo_date",
		"invoice_confirm_surname",
		"invoice_confirm_telephone",
		"invoice_confirm_currency",
		"invoice_receiver_firstname",
		"invoice_receiver_surname",
		"invoice_receiver_email",
		"invoice_receiver_telephone",
		"enabled",
		"created_at",
		"modified_at",
		"created_by",
		"modified_by",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *Customer, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "contract_id" == columnName {
			if point {
				return &ins.ContractId
			}
			return ins.ContractId
		}
		if "full_name" == columnName {
			if point {
				return &ins.FullName
			}
			return ins.FullName
		}
		if "project_id" == columnName {
			if point {
				return &ins.ProjectId
			}
			return ins.ProjectId
		}
		if "legal_status" == columnName {
			if point {
				return &ins.LegalStatus
			}
			return ins.LegalStatus
		}
		if "customer_status" == columnName {
			if point {
				return &ins.CustomerStatus
			}
			return ins.CustomerStatus
		}
		if "tax_id" == columnName {
			if point {
				return &ins.TaxId
			}
			return ins.TaxId
		}
		if "field_of_business" == columnName {
			if point {
				return &ins.FieldOfBusiness
			}
			return ins.FieldOfBusiness
		}
		if "office_country" == columnName {
			if point {
				return &ins.OfficeCountry
			}
			return ins.OfficeCountry
		}
		if "office_city" == columnName {
			if point {
				return &ins.OfficeCity
			}
			return ins.OfficeCity
		}
		if "office_street" == columnName {
			if point {
				return &ins.OfficeStreet
			}
			return ins.OfficeStreet
		}
		if "office_region" == columnName {
			if point {
				return &ins.OfficeRegion
			}
			return ins.OfficeRegion
		}
		if "office_postal_code" == columnName {
			if point {
				return &ins.OfficePostalCode
			}
			return ins.OfficePostalCode
		}
		if "office_no" == columnName {
			if point {
				return &ins.OfficeNo
			}
			return ins.OfficeNo
		}
		if "address_country" == columnName {
			if point {
				return &ins.AddressCountry
			}
			return ins.AddressCountry
		}
		if "address_city" == columnName {
			if point {
				return &ins.AddressCity
			}
			return ins.AddressCity
		}
		if "address_street" == columnName {
			if point {
				return &ins.AddressStreet
			}
			return ins.AddressStreet
		}
		if "address_telephone" == columnName {
			if point {
				return &ins.AddressTelephone
			}
			return ins.AddressTelephone
		}
		if "address_region" == columnName {
			if point {
				return &ins.AddressRegion
			}
			return ins.AddressRegion
		}
		if "address_postal_code" == columnName {
			if point {
				return &ins.AddressPostalCode
			}
			return ins.AddressPostalCode
		}
		if "address_no" == columnName {
			if point {
				return &ins.AddressNo
			}
			return ins.AddressNo
		}
		if "address_fax" == columnName {
			if point {
				return &ins.AddressFax
			}
			return ins.AddressFax
		}
		if "contact_full_Name" == columnName {
			if point {
				return &ins.ContactFullName
			}
			return ins.ContactFullName
		}
		if "contact_surname" == columnName {
			if point {
				return &ins.ContactSurname
			}
			return ins.ContactSurname
		}
		if "contact_telephone" == columnName {
			if point {
				return &ins.ContactTelephone
			}
			return ins.ContactTelephone
		}
		if "contact_firstname" == columnName {
			if point {
				return &ins.ContactFirstname
			}
			return ins.ContactFirstname
		}
		if "contact_email" == columnName {
			if point {
				return &ins.ContactEmail
			}
			return ins.ContactEmail
		}
		if "contact_position" == columnName {
			if point {
				return &ins.ContactPosition
			}
			return ins.ContactPosition
		}
		if "bank_name" == columnName {
			if point {
				return &ins.BankName
			}
			return ins.BankName
		}
		if "bank_account_name" == columnName {
			if point {
				return &ins.BankAccountName
			}
			return ins.BankAccountName
		}
		if "bankgiro" == columnName {
			if point {
				return &ins.Bankgiro
			}
			return ins.Bankgiro
		}
		if "iban" == columnName {
			if point {
				return &ins.Iban
			}
			return ins.Iban
		}
		if "bank_address" == columnName {
			if point {
				return &ins.BankAddress
			}
			return ins.BankAddress
		}
		if "bank_account_no" == columnName {
			if point {
				return &ins.BankAccountNo
			}
			return ins.BankAccountNo
		}
		if "bic_swift" == columnName {
			if point {
				return &ins.BicSwift
			}
			return ins.BicSwift
		}
		if "bank_currency" == columnName {
			if point {
				return &ins.BankCurrency
			}
			return ins.BankCurrency
		}
		if "invoice_confirm_firstname" == columnName {
			if point {
				return &ins.InvoiceConfirmFirstname
			}
			return ins.InvoiceConfirmFirstname
		}
		if "invoice_confirm_email" == columnName {
			if point {
				return &ins.InvoiceConfirmEmail
			}
			return ins.InvoiceConfirmEmail
		}
		if "invoice_confirm_duo_date" == columnName {
			if point {
				return &ins.InvoiceConfirmDuoDate
			}
			return ins.InvoiceConfirmDuoDate
		}
		if "invoice_confirm_surname" == columnName {
			if point {
				return &ins.InvoiceConfirmSurname
			}
			return ins.InvoiceConfirmSurname
		}
		if "invoice_confirm_telephone" == columnName {
			if point {
				return &ins.InvoiceConfirmTelephone
			}
			return ins.InvoiceConfirmTelephone
		}
		if "invoice_confirm_currency" == columnName {
			if point {
				return &ins.InvoiceConfirmCurrency
			}
			return ins.InvoiceConfirmCurrency
		}
		if "invoice_receiver_firstname" == columnName {
			if point {
				return &ins.InvoiceReceiverFirstname
			}
			return ins.InvoiceReceiverFirstname
		}
		if "invoice_receiver_surname" == columnName {
			if point {
				return &ins.InvoiceReceiverSurname
			}
			return ins.InvoiceReceiverSurname
		}
		if "invoice_receiver_email" == columnName {
			if point {
				return &ins.InvoiceReceiverEmail
			}
			return ins.InvoiceReceiverEmail
		}
		if "invoice_receiver_telephone" == columnName {
			if point {
				return &ins.InvoiceReceiverTelephone
			}
			return ins.InvoiceReceiverTelephone
		}
		if "enabled" == columnName {
			if point {
				return &ins.Enabled
			}
			return ins.Enabled
		}
		if "created_at" == columnName {
			if point {
				return &ins.CreatedAt
			}
			return ins.CreatedAt
		}
		if "modified_at" == columnName {
			if point {
				return &ins.ModifiedAt
			}
			return ins.ModifiedAt
		}
		if "created_by" == columnName {
			if point {
				return &ins.CreatedBy
			}
			return ins.CreatedBy
		}
		if "modified_by" == columnName {
			if point {
				return &ins.ModifiedBy
			}
			return ins.ModifiedBy
		}
		return nil
	},
}

type Customer struct {
	Id                       int64
	ContractId               ttypes.NilableString
	FullName                 string
	ProjectId                ttypes.NilableString
	LegalStatus              ttypes.NilableString
	CustomerStatus           ttypes.NilableString
	TaxId                    ttypes.NilableString
	FieldOfBusiness          ttypes.NilableString
	OfficeCountry            ttypes.NilableString
	OfficeCity               ttypes.NilableString
	OfficeStreet             ttypes.NilableString
	OfficeRegion             ttypes.NilableString
	OfficePostalCode         ttypes.NilableString
	OfficeNo                 ttypes.NilableString
	AddressCountry           ttypes.NilableString
	AddressCity              ttypes.NilableString
	AddressStreet            ttypes.NilableString
	AddressTelephone         ttypes.NilableString
	AddressRegion            ttypes.NilableString
	AddressPostalCode        ttypes.NilableString
	AddressNo                ttypes.NilableString
	AddressFax               ttypes.NilableString
	ContactFullName          ttypes.NilableString
	ContactSurname           ttypes.NilableString
	ContactTelephone         ttypes.NilableString
	ContactFirstname         ttypes.NilableString
	ContactEmail             ttypes.NilableString
	ContactPosition          ttypes.NilableString
	BankName                 ttypes.NilableString
	BankAccountName          ttypes.NilableString
	Bankgiro                 ttypes.NilableString
	Iban                     ttypes.NilableString
	BankAddress              ttypes.NilableString
	BankAccountNo            ttypes.NilableString
	BicSwift                 ttypes.NilableString
	BankCurrency             ttypes.NilableString
	InvoiceConfirmFirstname  ttypes.NilableString
	InvoiceConfirmEmail      ttypes.NilableString
	InvoiceConfirmDuoDate    ttypes.NilableString
	InvoiceConfirmSurname    ttypes.NilableString
	InvoiceConfirmTelephone  ttypes.NilableString
	InvoiceConfirmCurrency   ttypes.NilableString
	InvoiceReceiverFirstname ttypes.NilableString
	InvoiceReceiverSurname   ttypes.NilableString
	InvoiceReceiverEmail     ttypes.NilableString
	InvoiceReceiverTelephone ttypes.NilableString
	Enabled                  bool
	CreatedAt                ttypes.NilableDatetime
	ModifiedAt               ttypes.NilableDatetime
	CreatedBy                int64
	ModifiedBy               int64
}

func (c *Customer) ToDomainInfo() *model.DomainInfo {
	return &model.DomainInfo{
		DomainId:   c.Id,
		DomainName: c.FullName,
		Enabled:    c.Enabled,
	}
}
