package test

import (
	dgctx "dghire.com/libs/go-common/context"
	"github.com/google/uuid"
	"go-uc/internal/tlog"
)

func NewBusinessUserCTX() *tlog.Ctx {
	return tlog.NewCtx(&dgctx.DgContext{
		TraceId:  uuid.NewString(),
		UserId:   24,
		Role:     "GROUP_MANAGER",
		Domain:   "BUSINESS",
		DomainId: 1,
		GroupId:  1,
	})
}

func NewPlatformUserCTX() *tlog.Ctx {
	return tlog.NewCtx(&dgctx.DgContext{
		TraceId: uuid.NewString(),
		UserId:  25,
		Role:    "MANAGER",
		Domain:  "PLATFORM",
	})
}

func NewSupplierUserCTX() *tlog.Ctx {
	return tlog.NewCtx(&dgctx.DgContext{
		TraceId:  uuid.NewString(),
		UserId:   29,
		Role:     "SUPER_MANAGER",
		Domain:   "SUPPLIER",
		DomainId: 1,
	})
}

func NewEmployeeCTX() *tlog.Ctx {
	return tlog.NewCtx(&dgctx.DgContext{
		Domain:  "EMPLOYEE",
		TraceId: uuid.NewString(),
		UserId:  40,
	})
}

func NewBlankCTX() *tlog.Ctx {
	return tlog.NewCtx(&dgctx.DgContext{})
}
