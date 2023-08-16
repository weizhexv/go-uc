package test

import (
	"fmt"
	"github.com/rolandhe/daog"
	"github.com/rolandhe/daog/ttypes"
	txrequest "github.com/rolandhe/daog/tx"
	"go-uc/internal/dal"
	"go-uc/internal/repo"
	"go-uc/internal/store"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	context, err := daog.NewTransContext(store.DB(), txrequest.RequestWrite, "assjjjjjj")
	if err != nil {
		return
	}

	defer func() {
		context.Complete(err)
	}()

	n := &dal.Nonce{
		Id:        1,
		Identity:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaassssssssssssssJJJJJJJJJJJJJ",
		Nonce:     "",
		Status:    0,
		CreatedBy: 0,
		ExpireAt:  ttypes.NormalDatetime{},
		CreatedAt: ttypes.NormalDatetime{},
	}
	_, err = repo.Nonce.Update(context, n)
	if err != nil {
		fmt.Println("error ::::", err)
		return
	}
}

func TestExpiresAt(t *testing.T) {
	a := time.Now().AddDate(0, 0, 30)
	fmt.Println(time.Now())
	fmt.Println(a)
}
