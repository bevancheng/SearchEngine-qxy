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

type UserSignUpService struct {
	ctx context.Context
}

func NewUserSignUpService(ctx context.Context) *UserSignUpService {
	return &UserSignUpService{ctx: ctx}
}

func (usu *UserSignUpService) SignUp(req *user.UserSignUpRequest) error {
	users, err := db.QueryUser(usu.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(usu.ctx, []*db.User{{
		UserName: req.Username,
		PassWord: password,
	}})
}
