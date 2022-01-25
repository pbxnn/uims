package biz

import (
	"context"
	"errors"

	umsApi "uims/api/ums/api"
	umsService "uims/api/ums/rpc"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       int64
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*umsService.CreateUserReply, error)
	GetUser(ctx context.Context, id int64) (*umsService.GetUserReply, error)
	GetUserByUsername(ctx context.Context, username string) (*umsService.GetUserByUsernameReply, error)
	Save(ctx context.Context, u *User) (*umsService.SaveUserReply, error)
	ListUser(ctx context.Context, u *User) (*umsService.ListUserReply, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Save(ctx context.Context, in *umsApi.SaveUserReq) (*umsService.SaveUserReply, error) {
	user := &User{
		Username: in.Username,
	}
	data, err := uc.repo.Save(ctx, user)
	if err != nil {
		// todo: handle error
		return nil, err
	}
	return data, nil
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*umsService.CreateUserReply, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*umsService.GetUserReply, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, in *umsApi.GetUserByUsernameReq) (*umsService.GetUserByUsernameReply, error) {
	user, err := uc.repo.GetUserByUsername(ctx, in.Username)
	if err != nil {
		//todo: handle error
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) ListUser(ctx context.Context, u *User) (*umsService.ListUserReply, error) {
	out, err := uc.repo.ListUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}
