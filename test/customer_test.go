package test

import (
	dgctx "dghire.com/libs/go-common/context"
	"encoding/json"
	"fmt"
	"go-uc/internal/model"
	"go-uc/internal/service"
	"go-uc/internal/tlog"
	"testing"
)

func TestCustomerService(t *testing.T) {
	c := tlog.NewCtx(&dgctx.DgContext{TraceId: "aaa"})
	q := &model.CustomerQuery{
		Id:               0,
		FullName:         "",
		CheckEnabled:     true,
		Enabled:          false,
		ProjectId:        "",
		ContractFullName: "",
		OfficeCountry:    "",
		PageNo:           0,
		PageSize:         0,
	}
	ret, err := service.Customer.Query(c, q)
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := json.Marshal(ret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s", marshal)
}
