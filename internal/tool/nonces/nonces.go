package nonces

import (
	dgerr "dghire.com/libs/go-common/enums/error"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/errs"
	"strconv"
	"strings"
	"time"
)

func New(ctx *tlog.Ctx, identity string) (string, error) {
	if len(identity) == 0 {
		return "", dgerr.ARGUMENT_NOT_VALID
	}
	nc := dal.NewNonce(ctx.UserId, identity)

	err := daog.AutoTrans(func() (*daog.TransContext, error) {
		return daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	}, func(tc *daog.TransContext) error {
		_, err := repo.Nonce.Insert(tc, nc)
		return err
	})
	if err != nil {
		ctx.Errorf("create nonce err: %v", err)
		return "", err
	}

	return strconv.FormatInt(nc.Id, 10) + "_" + nc.Nonce, nil
}

func Check(ctx *tlog.Ctx, nonce string) error {
	if len(nonce) == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	id, raw, err := parseIdAndRaw(nonce)
	if err != nil {
		ctx.Error("parse id and raw err:", err)
		return dgerr.ARGUMENT_NOT_VALID
	}

	return daog.AutoTrans(func() (*daog.TransContext, error) {
		return daog.NewTransContext(store.DB(), txrequest.RequestReadonly, ctx.TraceId)
	}, func(tc *daog.TransContext) error {
		po, err := repo.Nonce.GetById(tc, id)
		if err != nil {
			return err
		}
		if po == nil {
			return errs.LinkExpired
		}
		if po.Nonce != raw || po.Status != 0 || time.Now().After(time.Time(po.ExpireAt)) {
			return errs.LinkExpired
		}
		return nil
	})
}

func parseIdAndRaw(nonce string) (int64, string, error) {
	sl := strings.Split(nonce, "_")
	id, err := strconv.ParseInt(sl[0], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return id, sl[1], nil
}

func Disable(ctx *tlog.Ctx, nonce string) error {
	if len(nonce) == 0 {
		return dgerr.ARGUMENT_NOT_VALID
	}
	id, raw, err := parseIdAndRaw(nonce)
	if err != nil {
		ctx.Error("parse id and raw err:", err)
		return dgerr.ARGUMENT_NOT_VALID
	}

	return daog.AutoTrans(func() (*daog.TransContext, error) {
		return daog.NewTransContext(store.DB(), txrequest.RequestWrite, ctx.TraceId)
	}, func(tc *daog.TransContext) error {
		rows, err := repo.Nonce.UpdateByModifier(tc, daog.NewModifier().Add(dal.NonceFields.Status, 1),
			daog.NewMatcher().Eq(dal.NonceFields.Id, id).Eq(dal.NonceFields.Nonce, raw))
		if err != nil {
			return err
		}
		if rows <= 0 {
			return dgerr.LOGIN_ERROR
		}
		return nil
	})
}
