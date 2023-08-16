package test

import (
	"fmt"
	"go-uc/internal/service"
	"go-uc/pkg/var/domains"
	"testing"
)

func TestDomains(t *testing.T) {
	//test Equal() func
	ok := domains.Equal("business", domains.Business)
	if !ok {
		t.Fatal("first")
	}
	ok = domains.Equal("biz", domains.Business)
	if ok {
		t.Fatal("second")
	}

	//test Parse() func
	dm := domains.Parse("Business")
	fmt.Println("dm is ", dm)
	if len(dm) == 0 {
		t.Fatal("third")
	}

	dm = domains.Parse("not exist")
	fmt.Println("not exist ", dm)
	if len(dm) != 0 {
		t.Fatal("Fourth")
	}
}

func TestGetDomainInfo(t *testing.T) {
	ctx := NewPlatformUserCTX()
	info, err := service.DomainService.GetDomainInfo(ctx, domains.Platform, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("info is ", info)
}
