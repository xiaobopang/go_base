package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"go-base/internal/model"

	v1 "go-base/api/v1"
	"go-base/internal/codes"
	"go-base/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Get(ctx context.Context, req *v1.UserGetReq) (*v1.UserGetRes, error) {
	userInfo, err := service.User().Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if userInfo == nil {
		return nil, gerror.WrapCode(codes.CodeNotFound, fmt.Errorf("id '%d'", req.Id))
	}

	return &v1.UserGetRes{
		Id:        userInfo.Id,
		Nickname:  userInfo.Nickname,
		Age:       userInfo.Age,
		Gender:    userInfo.Gender,
		Status:    userInfo.Status,
		CreatedAt: userInfo.CreatedAt,
		UpdatedAt: userInfo.UpdatedAt,
	}, nil
}

func (c *cUser) Create(ctx context.Context, req *v1.UserCreateReq) (*v1.UserCreateRes, error) {
	data := model.UserCreateInput{
		Nickname: req.Nickname,
		Age:      req.Age,
		Gender:   req.Gender,
	}

	res, err := service.User().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.UserCreateRes{ID: res.ID}, err
}

func (c *cUser) GetList(ctx context.Context, req *v1.UserListReq) (*v1.UserListRes, error) {
	res, err := service.User().List(ctx, model.UserListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := res.List
	if len(list) == 0 {
		// 避免返回的空数组为null
		list = []model.UserListOutputItem{}
	}

	return &v1.UserListRes{
		CommonPaginationRes: v1.CommonPaginationRes{
			Total:    res.Total,
			PageNum:  res.PageNum,
			PageSize: res.PageSize,
		},
		List: list,
	}, nil
}
