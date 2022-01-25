package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"uims/api/orgms/api"
	"uims/api/orgms/rpc"
)

type CompanyDO struct {
	c   rpc.CompanyClient
	log *log.Helper
}

func NewCompanyDO(c rpc.CompanyClient, logger log.Logger) *CompanyDO {
	return &CompanyDO{
		c:   c,
		log: log.NewHelper(log.With(logger, "module", "biz/company")),
	}
}

func (do *CompanyDO) CreateCompany(ctx context.Context, apiReq *api.CreateCompanyReq) (*api.CreateCompanyReply, error) {
	rpcReq := &rpc.CreateCompanyReq{}
	rpcReply := &rpc.CreateCompanyReply{}
	apiReply := &api.CreateCompanyReply{}
	err := copier.Copy(rpcReq, apiReq)
	if err != nil {
		return nil, err
	}
	rpcReply, err = do.c.CreateCompany(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(apiReply, rpcReply)
	if err != nil {
		return nil, err
	}

	return apiReply, nil
}
