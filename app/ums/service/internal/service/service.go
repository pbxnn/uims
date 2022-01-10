package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	umsApi "uims/api/ums/service"
	"uims/app/ums/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	umsApi.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}
