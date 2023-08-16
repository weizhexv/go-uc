package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-uc/internal/service"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/roles"
	"golang.org/x/exp/slices"
	"testing"
)

func TestRoles(t *testing.T) {
	r := roles.OfId(4)
	if r == nil || r != roles.SuperManager {
		t.Fatal("first")
	}

	r = roles.OfId(11)
	if r != nil {
		t.Fatal("second")
	}

	r = roles.OfName("hr")
	if r == nil || r != roles.HR {
		t.Fatal("third")
	}

	r = roles.OfName("not exist")
	if r != nil {
		t.Fatal("fourth")
	}
}

func TestRoleService(t *testing.T) {
	r := service.Role.FindByDomain(domains.Business)
	fmt.Printf("roles %+v", r)
	assert.True(t, slices.Contains(r, roles.SuperManager))
	assert.True(t, slices.Contains(r, roles.Finance))
}

func TestHasSuperManager(t *testing.T) {
	has, err := service.Role.HasSuperManager(NewPlatformUserCTX(), domains.Business, 1)
	assert.True(t, err == nil)
	assert.True(t, has)
	has, err = service.Role.HasSuperManager(NewPlatformUserCTX(), domains.Business, 2)
	assert.True(t, err == nil)
	assert.False(t, has)
}
