package biz

import (
	"context"
	"errors"

	umsService "uims/api/ums/rpc"
	"uims/app/ums/rpc/internal/data/dao"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	UId      int64
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*dao.UmsUser, error)
	GetUser(ctx context.Context, id int64) (*dao.UmsUser, error)
	GetUserByUsername(ctx context.Context, username string) (*dao.UmsUser, error)
	Save(ctx context.Context, u *User) (*dao.UmsUser, error)
	ListUser(ctx context.Context, u *User) ([]*dao.UmsUser, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Save(ctx context.Context, in *umsService.SaveUserReq) (*dao.UmsUser, error) {
	user := &User{
		Username: in.Username,
	}
	data, err := uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*dao.UmsUser, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*dao.UmsUser, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, in *umsService.GetUserByUsernameReq) (*dao.UmsUser, error) {
	user, err := uc.repo.GetUserByUsername(ctx, in.Username)
	if err != nil {
		//todo: handle error
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) ListUser(ctx context.Context, u *User) ([]*dao.UmsUser, error) {
	out, err := uc.repo.ListUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}
