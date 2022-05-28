package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"userser/cmd/user/dal/db"
	"userser/cmd/user/kitex_gen/user"
	"userser/pkg/errno"
)

type UserLogInService struct {
	ctx context.Context
}

func NewUserLogInService(ctx context.Context) *UserLogInService {
	return &UserLogInService{ctx: ctx}
}

func (uli *UserLogInService) LogIn(req *user.UserLogInRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	username := req.Username
	users, err := db.QueryUser(uli.ctx, username)
	if err != nil {
		return 0, nil
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.PassWord != password {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}
