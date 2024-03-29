// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-base/api/v1"
	"go-base/internal/model"
	"go-base/internal/model/entity"
)

type (
	IUser interface {
		Get(ctx context.Context, req *v1.UserGetReq) (*entity.User, error)
		Update(ctx context.Context, in model.UserUpdateInput) error
		Create(ctx context.Context, in model.UserCreateInput) (*model.UserCreateOutput, error)
		List(ctx context.Context, in model.UserListInput) (*model.UserListOutput, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
