package dal

import (
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"time"
)

var VerificationDecisionFields = struct {
	Id        string
	Body      string
	CreatedAt string
}{
	"id",
	"body",
	"created_at",
}

var VerificationDecisionMeta = &daog.TableMeta[VerificationDecision]{
	Table: "verification_decision",
	Columns: []string{
		"id",
		"body",
		"created_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *VerificationDecision, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "body" == columnName {
			if point {
				return &ins.Body
			}
			return ins.Body
		}
		if "created_at" == columnName {
			if point {
				return &ins.CreatedAt
			}
			return ins.CreatedAt
		}

		return nil
	},
}

var VerificationDecisionDao daog.QuickDao[VerificationDecision] = &struct {
	daog.QuickDao[VerificationDecision]
}{
	daog.NewBaseQuickDao(VerificationDecisionMeta),
}

type VerificationDecision struct {
	Id        int64
	Body      string
	CreatedAt ttypes.NormalDatetime
}

func SaveVerificationDecision(tc *daog.TransContext, body string) error {
	_, err := VerificationDecisionDao.Insert(tc, &VerificationDecision{
		Body:      body,
		CreatedAt: ttypes.NormalDatetime(time.Now()),
	})

	return err
}
