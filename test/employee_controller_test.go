package test

import (
	dgctx "dghire.com/libs/go-common/context"
	dghttp "dghire.com/libs/go-httpclient"
	"dghire.com/libs/go-web/wrapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-uc/internal/proxy/veriff"
	"go-uc/internal/service"
	"go-uc/internal/tool/nonces"
	"go-uc/pkg/req"
	"go-uc/pkg/vo"
	"go-uc/web/controller"
	"testing"
)

func TestEmployeeSignUp(t *testing.T) {
	email := "4567@qq.com"

	nc, err := nonces.New(NewSupplierUserCTX(), email)
	assert.Nil(t, err)
	assert.True(t, len(nc) > 0)
	fmt.Printf("nonce is %s", nc)

	gc := new(gin.Context)
	ctx := NewBlankCTX().DgContext
	reqObj := &req.SignUpParams{
		Name:       "伟雇员",
		Email:      email,
		Password:   "12345678",
		Country:    "SWE",
		TaxCountry: "SWE",
		Birthday:   "1992-08-07",
		Passport:   uuid.NewString(),
		Mobile:     "13521511968",
		Address:    "顺德大良新区",
		PostalCode: "10086",
		Nonce:      nc,
		CardType:   "passport",
	}

	f := controller.Employee.SignUp()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestEmployeeAccountId(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewEmployeeCTX().DgContext
	reqObj := &req.UpdateEmployeeParams{
		AccountId: 101,
	}

	f := controller.Employee.Update()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestEmployeeUpdate(t *testing.T) {
	var banks []*vo.BankAccount
	acc1 := &vo.BankAccount{
		BankName:  "工商",
		AccountNo: "伟公章",
		OwnerName: "伟小牛渐渐",
	}
	acc2 := &vo.BankAccount{
		BankName:  "招商",
		AccountNo: "伟招商",
		OwnerName: "伟小妞",
	}
	banks = append(banks, acc1, acc2)

	var files []*vo.FgwFile
	f1 := &vo.FgwFile{
		FileId:   1,
		Filename: "jkqj",
		Format:   "pdf",
		Size:     100,
		Url:      "",
	}
	f2 := &vo.FgwFile{
		FileId:   2,
		Filename: "reta",
		Format:   "txt",
		Size:     100,
		Url:      "",
	}
	files = append(files, f1, f2)

	gc := new(gin.Context)
	ctx := NewEmployeeCTX().DgContext
	reqObj := &req.UpdateEmployeeParams{
		IdCard:       "11011010101010",
		Attachments:  files,
		BankAccounts: banks,
		AccountId:    100,
		CardType:     "passport",
	}

	f := controller.Employee.Update()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestEmployeeInfo(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewEmployeeCTX().DgContext
	reqObj := &wrapper.MapRequest{}
	f := controller.Employee.Info()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestEmployeeInfoByPlatform(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewPlatformUserCTX().DgContext
	reqObj := &wrapper.MapRequest{
		MP: map[string]any{},
	}
	reqObj.MP["employeeId"] = "38"

	f := controller.Employee.Info()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestSubmitVerify(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewEmployeeCTX().DgContext
	reqObj := &wrapper.MapRequest{MP: map[string]any{"verified": "false"}}
	f := controller.Employee.SubmitVerify()
	ret := f(gc, ctx, reqObj)
	assert.True(t, ret.Success)
	fmt.Printf("ret %v", ret)
}

func TestEmployeeExist(t *testing.T) {
	gc := new(gin.Context)
	ctx := NewBlankCTX().DgContext
	reqObj := &wrapper.MapRequest{MP: map[string]any{"email": "996997257@qq.com"}}
	f := controller.Employee.Exist()
	ret := f(gc, ctx, reqObj)

	assert.True(t, ret.Success)
	fmt.Printf("ret: %v", ret)
}

func TestVerificationDecision(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "xxxyyyzzz"}
	veriffCallbackBody := &service.VeriffDecisionBody{
		Status: "success",
		Verification: service.VeriffVerification{
			Id:     "12df6045-3846-3e45-946a-14fa6136d78b",
			Code:   9001,
			Status: "approved",
			Document: service.VeriffDocument{
				Type:   "ID_CARD",
				Number: "110113",
			},
		},
	}
	headers := map[string]string{
		"X-HMAC-SIGNATURE": veriff.HmacSHA256("12df6045-3846-3e45-946a-14fa6136d78b"),
	}
	dghttp.GlobalHttpClient.UseMonitor = false
	resp, err := dghttp.GlobalHttpClient.DoPostJson(ctx, "http://localhost:9999/uc/public/employee/verification-decision", veriffCallbackBody, headers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}
