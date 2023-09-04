package user

import (
	"context"
	"errors"
	"zxcblog/study_blog/pb/user"
	"zxcblog/study_blog/pkg/errcode"
)

type UserService struct {
}

func NewUserService() user.UserServer {
	return UserService{}
}

func (u UserService) Register(ctx context.Context, req *user.RegisterReq) (*user.UserAuthRes, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	//TODO implement me
	return nil, errcode.ToGrpcError(errcode.NewError(1234, "测试错误代码"))
}

func (u UserService) Login(ctx context.Context, req *user.LoginReq) (*user.UserAuthRes, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}
