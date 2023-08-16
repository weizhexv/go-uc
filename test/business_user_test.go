package test

import (
	"encoding/json"
	"fmt"
	"github.com/rolandhe/daog"
	txrequest "github.com/rolandhe/daog/tx"
	"github.com/stretchr/testify/assert"
	"go-uc/internal/dal"
	"go-uc/internal/model"
	"go-uc/internal/repo"
	"go-uc/internal/service"
	"go-uc/internal/store"
	"testing"
)

func TestBusinessUser(t *testing.T) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, "traceidxxx")
	assert.Nil(t, err)
	q := &model.UserQuery{
		CheckEnabled: true,
		Enabled:      true,
		PageNo:       1,
		PageSize:     100,
	}
	infos, err := service.BusinessUserImpl.Query(nil, tc, q)
	assert.Nil(t, err)
	j, err := json.Marshal(infos)
	assert.Nil(t, err)
	fmt.Printf("infos %s", j)
}

func TestDalBusinessUser(t *testing.T) {
	tc, err := daog.NewTransContext(store.DB(), txrequest.RequestReadonly, "traceidxxx")
	assert.Nil(t, err)
	mc := daog.NewMatcher().Eq(dal.BusinessUserFields.Enabled, false)
	ret, err := repo.BusinessUser.QueryListMatcher(tc, mc)
	assert.Nil(t, err)
	j, err := json.Marshal(ret)
	assert.Nil(t, err)
	fmt.Printf("ret: %s", j)
}
