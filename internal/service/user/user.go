package user

import (
	"context"
	"errors"
	"zxcblog/study_blog/pb/user"
)

type UserService struct {
}

func NewUserService() user.UserServer {
	return UserService{}
}

func (u UserService) Register(ctx context.Context, req *user.RegisterReq) (*user.UserAuthRes, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}

func (u UserService) Login(ctx context.Context, req *user.LoginReq) (*user.UserAuthRes, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}
