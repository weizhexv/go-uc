package service

import (
	"go-uc/internal/model"
	"go-uc/internal/tlog"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/var/errs"
)

type domainService struct {
}

var DomainService = new(domainService)

func (ds *domainService) MapDomainIdNames(c *tlog.Ctx, domain domains.Domain, domainIds []int64) (map[int64]string, error) {
	c.Infof("map domainId Names by domain: %v, domainIds: %v", domain, domainIds)

	domainNames := make(map[int64]string)
	if domains.Business.Is(domain) {
		customers, err := Customer.FindByIds(c, domainIds)
		if err != nil {
			return nil, err
		}
		for _, e := range customers {
			domainNames[e.Id] = e.FullName
		}
	} else if domains.Supplier.Is(domain) {
		suppliers, err := Supplier.FindByIds(c, domainIds)
		if err != nil {
			return nil, err
		}
		for _, e := range suppliers {
			domainNames[e.Id] = e.FullName
		}
	}

	c.Infof("map domainId Names ret %v", domainNames)
	return domainNames, nil
}

func (ds *domainService) GetDomainInfo(c *tlog.Ctx, domain domains.Domain, domainId int64) (*model.DomainInfo, error) {
	c.Infof("get domain info by domainId: %d", domainId)

	var dmInfo *model.DomainInfo
	if domain.Is(domains.Business) {
		cus, err := Customer.FindById(c, domainId)
		if err != nil {
			return nil, err
		}
		if cus == nil {
			return nil, errs.CompanyDisabled
		}
		dmInfo = cus.ToDomainInfo()
	} else if domain.Is(domains.Supplier) {
		spl, err := Supplier.FindById(c, domainId)
		if err != nil {
			return nil, err
		}
		if spl == nil {
			return nil, errs.CompanyDisabled
		}
		dmInfo = spl.ToDomainInfo()
	} else {
		dmInfo = &model.DomainInfo{
			DomainId:   0,
			DomainName: "",
			Enabled:    true,
		}
	}

	c.Infof("get domain info ret: %v", dmInfo)
	return dmInfo, nil
}
