package service

import (
	"errors"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/store"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/emails"
)

type emailService struct {
}

var Email = new(emailService)

func (e *emailService) SendRegisterEmail(c *tlog.Ctx, tc *daog.TransContext, uid int64) error {
	c.Infof("send register email by uid: %d", uid)
	detail, err := UserImpls.FindDetailByUid(c, tc, uid)
	if err != nil {
		c.Errorf("find detail by uid error: %v", err)
		return err
	}
	return emails.SendRegisterEmail(c, detail)
}

func (e *emailService) SendResetPasswordEmail(c *tlog.Ctx, uid int64) error {
	c.Infof("send reset password email by uid %d", uid)
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, c.TraceId)
	if err != nil {
		c.Errorf("new trans error: %v", err)
		return err
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

	detail, err := UserImpls.FindDetailByUid(c, tc, uid)
	if err != nil {
		return err
	}
	err = emails.SendResetPasswordEmail(c, detail)
	c.Errorf("send reset password error: %v", err)
	return err
}
