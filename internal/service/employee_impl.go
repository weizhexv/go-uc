package service

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	dgi18n "dghire.com/libs/go-i18n"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/convs"
	"go-uc/internal/dal"
	"go-uc/internal/enum"
	"go-uc/internal/model"
	"go-uc/internal/proxy/veriff"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
	"go-uc/pkg/vo"
	"strconv"
	"strings"
)

var EmployeeImpl = new(employee)

type employee struct{}

func (e *employee) domain() domains.Domain {
	return domains.Employee
}

func (e *employee) Query(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) ([]*vo.UserInfo, error) {
	c.Infof("employee query by %v", query)
	page, err := repo.Employee.QueryListMatcher(tc, e.queryToMatcher(query), daog.NewDescOrder(dal.EmployeeFields.Id))
	if err != nil {
		c.Errorf("employee query error: %v", err)
		return nil, err
	}
	return e.toInfos(page), nil
}

func (e *employee) toInfos(employees []*dal.Employee) []*vo.UserInfo {
	if len(employees) == 0 {
		return []*vo.UserInfo{}
	}
	infos := make([]*vo.UserInfo, len(employees))
	i := 0
	for _, ele := range employees {
		infos[i] = ele.ToInfo()
		i++
	}
	return infos
}

func (e *employee) queryToMatcher(query *model.UserQuery) daog.Matcher {
	mc := daog.NewMatcher()
	if len(query.Name) > 0 {
		mc.Like(dal.EmployeeFields.Name, query.Name, daog.LikeStyleAll)
	}
	if len(query.Mobile) > 0 {
		mc.Like(dal.EmployeeFields.Mobile, query.Mobile, daog.LikeStyleAll)
	}
	if len(query.Email) > 0 {
		mc.Like(dal.EmployeeFields.Email, query.Email, daog.LikeStyleAll)
	}
	if query.CheckEnabled {
		mc.Eq(dal.EmployeeFields.Enabled, query.Enabled)
	}
	return mc
}

func (e *employee) Count(c *tlog.Ctx, tc *daog.TransContext, query *model.UserQuery) (int64, error) {
	c.Infof("employee count by query %v", query)
	return repo.Employee.Count(tc, e.queryToMatcher(query))
}

func (e *employee) Insert(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("employee insert by %v", upsert)
	exist, err := repo.Employee.FindByUid(tc, upsert.Uid)
	if err != nil {
		c.Errorf("employee find by uid error: %v", err)
		return 0, err
	}
	if exist != nil {
		return 0, errs.UserExist
	}

	po, err := convs.NewEmployeeConv(upsert).ToPO(c.DgContext)
	if err != nil {
		c.Errorf("employee convs to po error: %v", err)
		return 0, err
	}

	_, err = repo.Employee.Insert(tc, po)
	if err != nil {
		c.Errorf("employee insert error: %v", err)
		return 0, err
	}
	return po.Id, nil
}

func (e *employee) Update(c *tlog.Ctx, tc *daog.TransContext, upsert *model.UpsertUser) (int64, error) {
	c.Infof("employee update by %v", upsert)
	po, err := repo.Employee.FindByUid(tc, upsert.Uid)
	if err != nil {
		c.Errorf("employee find by uid error: %v", err)
		return 0, err
	}
	if po == nil {
		return 0, dgerr.USER_NOT_EXISTS
	}

	err = convs.NewEmployeeConv(upsert).UpdatePO(c.DgContext, po)
	if err != nil {
		c.Errorf("employee convs to po error: %v", err)
		return 0, err
	}

	c.Infof("update employee po: %v", po)
	_, err = repo.Employee.Update(tc, po)
	if err != nil {
		c.Errorf("update employee error: %v", err)
		return 0, err
	}
	return po.Uid, nil
}

func (e *employee) ExistByEmail(c *tlog.Ctx, email string) (bool, error) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("new trans error: %v", err)
		return false, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Errorln("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	po, err := e.FindByEmail(c, tc, email)
	if err != nil {
		c.Errorf("employee find by email error: %v", err)
		return false, err
	}

	return po != nil, nil
}

func (e *employee) FindByEmail(c *tlog.Ctx, tc *daog.TransContext, email string) (*vo.UserInfo, error) {
	c.Infof("employee find by email: %s", email)
	po, err := repo.Employee.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.EmployeeFields.Email, email))
	if err != nil {
		c.Errorf("employee find one error: %v", err)
		return nil, err
	}
	return po.ToInfo(), nil
}

func (e *employee) FindByUid(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*vo.UserInfo, error) {
	c.Infof("employee find by uid: %d", uid)
	if po, err := repo.Employee.QueryOneMatcher(tc, daog.NewMatcher().Eq(dal.EmployeeFields.Uid, uid)); err != nil {
		c.Errorf("employee query one error: %v", err)
		return nil, err
	} else {
		return po.ToInfo(), nil
	}
}

func (e *employee) FindByUids(c *tlog.Ctx, tc *daog.TransContext, uids []int64) ([]*vo.UserInfo, error) {
	c.Infof("employee find by uids: %v", uids)
	pos, err := repo.Employee.QueryListMatcher(tc, daog.NewMatcher().In(dal.EmployeeFields.Uid, daog.ConvertToAnySlice(uids)))
	if err != nil {
		c.Errorf("employee find by uids error: %v", err)
		return nil, err
	}
	return e.toInfos(pos), nil
}

func (e *employee) ResetPassword(c *tlog.Ctx, tc *daog.TransContext, hash string, uid int64) error {
	c.Infof("reset password by hash: %s, uid: %d", hash, uid)
	po, err := repo.Employee.FindByUid(tc, uid)
	if err != nil {
		c.Errorf("reset password error: %v", err)
		return err
	}
	if po == nil {
		return dgerr.USER_NOT_EXISTS
	}
	rows, err := repo.Employee.UpdateByModifier(tc, daog.NewModifier().Add(dal.EmployeeFields.Password, hash),
		daog.NewMatcher().Eq(dal.EmployeeFields.Uid, uid))
	if err != nil {
		c.Errorf("employee update by modifier error: %v", err)
		return err
	}
	c.Infof("reset password success rows: %d", rows)
	return nil
}

func (e *employee) FindByDomainId(c *tlog.Ctx, tc *daog.TransContext, domainId int64) ([]*vo.UserInfo, error) {
	c.Infof("employee find all")
	all, err := repo.Employee.GetAll(tc)
	if err != nil {
		c.Errorf("employee find all error: %v", err)
		return nil, err
	}
	return e.toInfos(all), nil
}

func (e *employee) MapDomainHasSuperManager(c *tlog.Ctx, tc *daog.TransContext, domainIds []int64) (map[int64]bool, error) {
	c.Errorf("employee MapDomainHasSuperManager not support")
	return nil, dgerr.ILLEGAL_OPERATION
}

func (e *employee) FindUserModel(c *tlog.Ctx, tc *daog.TransContext, uid int64) (*model.UserModel, error) {
	c.Infof("find employee model by uid: %d", uid)
	po, err := repo.Employee.FindByUid(tc, uid)
	if err != nil {
		c.Errorf("find employee by uid error: %v", err)
		return nil, err
	}
	return po.ToModel(), nil
}

func (e *employee) FindEmployeeInfo(c *tlog.Ctx, uid int64) (*vo.EmployeeInfo, error) {
	c.Infof("find employee po by uid: %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("new trans error: %v", err)
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			c.Error("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	po, err := repo.Employee.FindByUid(tc, uid)
	if err != nil {
		c.Errorf("find employee by uid error: %v", err)
	}
	if po == nil {
		return nil, dgerr.USER_NOT_EXISTS
	}

	info, err := po.ToEmployeeInfo()
	if err != nil {
		c.Errorf("employee to info error: %v", err)
		return nil, err
	}

	vs, err := dal.FindVerificationSessionByEmployeeId(tc, uid)
	if err != nil {
		return nil, err
	}

	var verificationStatus string
	var reason string
	if vs != nil {
		verificationStatus = vs.Status
		reason = vs.Reason.StringNilAsEmpty()
		if "id_number_not_matched" == reason {
			reason = dgi18n.MustGetMessage(c.Lang, "uc.id_number_not_matched")
		}
	} else {
		verificationStatus = enum.VERIFICATION_STATUS_UNVERIFIED
	}
	verificationStatusName := dgi18n.MustGetMessage(c.Lang, "uc.verification_status_"+verificationStatus)
	info.VerificationDecision = &vo.VerificationDecisionVo{
		Status:     verificationStatus,
		StatusName: verificationStatusName,
		Reason:     reason,
	}

	return info, nil
}

func (e *employee) SubmitVerify(ctx *tlog.Ctx, sessionId string, verified bool) error {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			fmt.Println("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	err = e.UpdateVerify(ctx, tc, verified)
	if err != nil {
		return err
	}

	employeeId := ctx.UserId

	vs, er := dal.FindVerificationSessionBySessionId(tc, sessionId)
	if er != nil {
		err = er
		return err
	}

	if vs == nil {
		err = dal.CreateVerificationSession(tc, sessionId, employeeId, enum.VERIFICATION_STATUS_UNVERIFIED)
		if err != nil {
			return err
		}
	} else {
		err = dal.UpdateVerificationSessionEmployeeIdBySessionId(tc, sessionId, employeeId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *employee) UpdateVerify(ctx *tlog.Ctx, tc *daog.TransContext, verified bool) error {
	md := daog.NewModifier().Add(dal.EmployeeFields.Verified, verified)
	mc := daog.NewMatcher().Eq(dal.EmployeeFields.Uid, ctx.UserId)
	_, err := repo.Employee.UpdateByModifier(tc, md, mc)
	if err != nil {
		return err
	}

	return nil
}

type VeriffEventBody struct {
	Id         string `json:"id"`
	Code       int    `json:"code"`
	VendorData string `json:"vendorData"`
}

type VeriffDecisionBody struct {
	Status       string             `json:"status"`
	Verification VeriffVerification `json:"verification"`
}

type VeriffVerification struct {
	Id         string         `json:"id"`
	Code       int            `json:"code"`
	Status     string         `json:"status"`
	Reason     string         `json:"reason"`
	Document   VeriffDocument `json:"document"`
	VendorData string         `json:"vendorData"`
}

type VeriffDocument struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type VeriffRiskLabel struct {
	Category string `json:"category"`
}

func (e *employee) VerificationEvent(ctx *tlog.Ctx, signature string, body []byte) error {
	ctx.Printf("verification event, signature: %s, body:%s", signature, string(body))
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			fmt.Println("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	veriffBody := &VeriffEventBody{}
	err = json.Unmarshal(body, veriffBody)
	if err != nil {
		ctx.Printf("unmarshal veriff body error: %v", err)
		return err
	}

	sessionId := veriffBody.Id

	if sessionId == "" {
		ctx.Printf("unmarshal veriff body fail: %v", veriffBody)
		return errors.New("unmarshal veriff body fail")
	}

	if signature != veriff.HmacSHA256(string(body)) {
		ctx.Printf("signature is not matched, signature: %s, sessionId: %s", signature, sessionId)
		return errors.New("signature is not matched")
	}

	vs, er := dal.FindVerificationSessionBySessionId(tc, sessionId)
	if er != nil {
		err = er
		return err
	}

	if vs == nil && veriffBody.Code == 7002 {
		var employeeId int64
		if strings.Trim(veriffBody.VendorData, " ") != "" {
			employeeId, er = strconv.ParseInt(veriffBody.VendorData, 10, 64)
			if er != nil {
				err = er
				return err
			}
		}

		err := dal.CreateVerificationSession(tc, sessionId, employeeId, enum.VERIFICATION_STATUS_APPROVING)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *employee) VerificationDecision(ctx *tlog.Ctx, signature string, body []byte) error {
	ctx.Printf("verification decistion, signature: %s, body:%s", signature, string(body))
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("panic:", r))
		}
		if err != nil {
			fmt.Println("catch error in defer: ", err)
		}
		tc.Complete(err)
	}()

	err = dal.SaveVerificationDecision(tc, string(body))
	if err != nil {
		return err
	}

	veriffBody := &VeriffDecisionBody{}
	err = json.Unmarshal(body, veriffBody)
	if err != nil {
		ctx.Printf("unmarshal veriff body error: %v", err)
		return err
	}

	sessionId := veriffBody.Verification.Id

	if veriffBody.Status == "" || sessionId == "" {
		ctx.Printf("unmarshal veriff body fail: %v", veriffBody)
		return errors.New("unmarshal veriff body fail")
	}

	if signature != veriff.HmacSHA256(string(body)) {
		ctx.Printf("signature is not matched, signature: %s, sessionId: %s", signature, sessionId)
		return errors.New("signature is not matched")
	}

	vs, er := dal.FindVerificationSessionBySessionId(tc, sessionId)
	if er != nil {
		err = er
		return err
	}

	var employeeId int64
	if strings.Trim(veriffBody.Verification.VendorData, " ") != "" {
		employeeId, er = strconv.ParseInt(veriffBody.Verification.VendorData, 10, 64)
		if er != nil {
			err = er
			return err
		}
	}

	if vs == nil {
		var status string
		if veriffBody.Verification.Status == enum.VERIFICATION_STATUS_APPROVED {
			status = enum.VERIFICATION_STATUS_APPROVED
		} else {
			status = enum.VERIFICATION_STATUS_REJECTED
		}
		err := dal.CreateVerificationSession(tc, sessionId, employeeId, status)
		if err != nil {
			return err
		}
	} else {
		verified := veriffBody.Verification.Status == enum.VERIFICATION_STATUS_APPROVED
		err = e.UpdateVerify(ctx, tc, verified)
		if err != nil {
			return err
		}

		var status string
		var reason string

		if verified {
			if vs.EmployeeId > 0 {
				employeeId = vs.EmployeeId
				emp, _ := repo.Employee.FindByUid(tc, employeeId)
				if emp != nil {
					if veriffBody.Verification.Document.Type == "PASSPORT" && veriffBody.Verification.Document.Number == emp.Passport.StringNilAsEmpty() {
						status = enum.VERIFICATION_STATUS_APPROVED
					} else if veriffBody.Verification.Document.Type == "ID_CARD" && veriffBody.Verification.Document.Number == emp.IdCard.StringNilAsEmpty() {
						status = enum.VERIFICATION_STATUS_APPROVED
					} else {
						status = enum.VERIFICATION_STATUS_REJECTED
						reason = "id_number_not_matched"
					}
				} else {
					status = enum.VERIFICATION_STATUS_APPROVED
				}
			} else {
				status = enum.VERIFICATION_STATUS_APPROVED
			}
		} else {
			status = enum.VERIFICATION_STATUS_REJECTED
			reason = veriffBody.Verification.Reason
		}

		err = dal.UpdateVerificationSessionStatusBySessionId(tc, sessionId, status, reason)
		if err != nil {
			return err
		}
	}

	return nil
}
