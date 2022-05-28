package main

import (
	"context"
	"userser/cmd/user/kitex_gen/user"
	"userser/cmd/user/pack"
	"userser/cmd/user/service"
	"userser/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserSignUp implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserSignUp(ctx context.Context, req *user.UserSignUpRequest) (resp *user.UserSignUpResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserSignUpResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	}

	err = service.NewUserSignUpService(ctx).SignUp(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil //not return userid and token yet
}

// UserLogIn implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogIn(ctx context.Context, req *user.UserLogInRequest) (resp *user.UserLogInResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLogInResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	}

	userid, err := service.NewUserLogInService(ctx).LogIn(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.UesrId = userid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UserLogOut implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogOut(ctx context.Context, req *user.UserLogOutRequest) (resp *user.UserLogOutResponse, err error) {
	// TODO: Your code here...
	return
}
