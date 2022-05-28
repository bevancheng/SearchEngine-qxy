package db

import (
	"context"
	"userser/pkg/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUser(ctx context.Context, users []*User) error {

	return DB.WithContext(ctx).Create(users).Error

}

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
