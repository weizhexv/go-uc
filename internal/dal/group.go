package dal

import (
	"encoding/json"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"go-uc/pkg/var/domains"
	"time"
)

var GroupFields = struct {
	Id          string
	Name        string
	Description string
	Domain      string
	DomainId    string
	CountryCode string
	Deleted     string
	Source      string
	CreatedBy   string
	CreatedAt   string
	ModifiedBy  string
	ModifiedAt  string
}{
	"id",
	"name",
	"description",
	"domain",
	"domain_id",
	"country_code",
	"deleted",
	"source",
	"created_by",
	"created_at",
	"modified_by",
	"modified_at",
}

var GroupMeta = &daog.TableMeta[Group]{
	Table: "`group`",
	Columns: []string{
		"id",
		"name",
		"description",
		"domain",
		"domain_id",
		"country_code",
		"deleted",
		"source",
		"created_by",
		"created_at",
		"modified_by",
		"modified_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *Group, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "name" == columnName {
			if point {
				return &ins.Name
			}
			return ins.Name
		}
		if "description" == columnName {
			if point {
				return &ins.Description
			}
			return ins.Description
		}
		if "domain" == columnName {
			if point {
				return &ins.Domain
			}
			return ins.Domain
		}
		if "domain_id" == columnName {
			if point {
				return &ins.DomainId
			}
			return ins.DomainId
		}
		if "country_code" == columnName {
			if point {
				return &ins.CountryCode
			}
			return ins.CountryCode
		}
		if "deleted" == columnName {
			if point {
				return &ins.Deleted
			}
			return ins.Deleted
		}
		if "source" == columnName {
			if point {
				return &ins.Source
			}
			return ins.Source
		}
		if "created_by" == columnName {
			if point {
				return &ins.CreatedBy
			}
			return ins.CreatedBy
		}
		if "created_at" == columnName {
			if point {
				return &ins.CreatedAt
			}
			return ins.CreatedAt
		}
		if "modified_by" == columnName {
			if point {
				return &ins.ModifiedBy
			}
			return ins.ModifiedBy
		}
		if "modified_at" == columnName {
			if point {
				return &ins.ModifiedAt
			}
			return ins.ModifiedAt
		}

		return nil
	},
}

type Group struct {
	Id          int64
	Name        string
	Description string
	Domain      string
	DomainId    int64
	CountryCode string
	Deleted     int8
	Source      int8
	CreatedBy   int64
	CreatedAt   ttypes.NormalDatetime
	ModifiedBy  int64
	ModifiedAt  ttypes.NormalDatetime
}

func (g *Group) String() string {
	js, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

const (
	grpSourceSystem = iota
	grpSourceSuperManager
)

func NewDefaultGroup(domain domains.Domain, domainId int64) *Group {
	return &Group{
		Name:        "默认团队",
		Description: "系统自动创建，可按实际使用情况修改",
		Domain:      domain.String(),
		DomainId:    domainId,
		CountryCode: "ALL",
		Deleted:     0,
		Source:      grpSourceSystem,
		CreatedBy:   0,
		CreatedAt:   ttypes.NormalDatetime(time.Now()),
		ModifiedBy:  0,
		ModifiedAt:  ttypes.NormalDatetime(time.Now()),
	}
}
