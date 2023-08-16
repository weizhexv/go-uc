package dal

import (
	"errors"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	"time"
)

var VerificationSessionFields = struct {
	Id         string
	SessionId  string
	EmployeeId string
	Status     string
	Reason     string
	CreatedAt  string
	ModifiedAt string
}{
	"id",
	"session_id",
	"employee_id",
	"status",
	"reason",
	"created_at",
	"modified_at",
}

var VerificationSessionMeta = &daog.TableMeta[VerificationSession]{
	Table: "verification_session",
	Columns: []string{
		"id",
		"session_id",
		"employee_id",
		"status",
		"reason",
		"created_at",
		"modified_at",
	},
	AutoColumn: "id",
	LookupFieldFunc: func(columnName string, ins *VerificationSession, point bool) any {
		if "id" == columnName {
			if point {
				return &ins.Id
			}
			return ins.Id
		}
		if "session_id" == columnName {
			if point {
				return &ins.SessionId
			}
			return ins.SessionId
		}
		if "employee_id" == columnName {
			if point {
				return &ins.EmployeeId
			}
			return ins.EmployeeId
		}
		if "status" == columnName {
			if point {
				return &ins.Status
			}
			return ins.Status
		}
		if "reason" == columnName {
			if point {
				return &ins.Reason
			}
			return ins.Reason
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

		return nil
	},
}

var VerificationSessionDao daog.QuickDao[VerificationSession] = &struct {
	daog.QuickDao[VerificationSession]
}{
	daog.NewBaseQuickDao(VerificationSessionMeta),
}

type VerificationSession struct {
	Id         int64
	SessionId  string
	EmployeeId int64
	Status     string
	Reason     ttypes.NilableString
	CreatedAt  ttypes.NormalDatetime
	ModifiedAt ttypes.NormalDatetime
}

func CreateVerificationSession(tc *daog.TransContext, sessionId string, employeeId int64, status string) error {
	session, err := VerificationSessionDao.QueryOneMatcher(tc, daog.NewMatcher().Eq(VerificationSessionFields.SessionId, sessionId))
	if err != nil {
		return err
	}
	if session != nil {
		return nil
	}

	now := ttypes.NormalDatetime(time.Now())
	_, err = VerificationSessionDao.Insert(tc, &VerificationSession{
		SessionId:  sessionId,
		EmployeeId: employeeId,
		Status:     status,
		CreatedAt:  now,
		ModifiedAt: now,
	})

	return err
}

func FindVerificationSessionBySessionId(tc *daog.TransContext, sessionId string) (*VerificationSession, error) {
	return VerificationSessionDao.QueryOneMatcher(tc, daog.NewMatcher().Eq(VerificationSessionFields.SessionId, sessionId))
}

func FindVerificationSessionByEmployeeId(tc *daog.TransContext, employeeId int64) (*VerificationSession, error) {
	matcher := daog.NewMatcher().Eq(VerificationSessionFields.EmployeeId, employeeId)
	order := daog.NewDescOrder(VerificationSessionFields.CreatedAt)
	list, err := VerificationSessionDao.QueryListMatcher(tc, matcher, order)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

func UpdateVerificationSessionStatusBySessionId(tc *daog.TransContext, sessionId string, status string, reason string) error {
	matcher := daog.NewMatcher().Eq(VerificationSessionFields.SessionId, sessionId)
	session, err := VerificationSessionDao.QueryOneMatcher(tc, matcher)
	if err != nil {
		return err
	}

	if session == nil {
		return errors.New("no record")
	}

	session.Status = status
	session.Reason = *ttypes.FromString(reason)
	session.ModifiedAt = ttypes.NormalDatetime(time.Now())

	_, err = VerificationSessionDao.Update(tc, session)
	if err != nil {
		return err
	}

	return nil
}

func UpdateVerificationSessionEmployeeIdBySessionId(tc *daog.TransContext, sessionId string, employeeId int64) error {
	matcher := daog.NewMatcher().Eq(VerificationSessionFields.SessionId, sessionId)
	modifier := daog.NewModifier().Add(VerificationSessionFields.EmployeeId, employeeId).Add(VerificationSessionFields.ModifiedAt, ttypes.NormalDatetime(time.Now()))

	_, err := VerificationSessionDao.UpdateByModifier(tc, modifier, matcher)
	return err
}
