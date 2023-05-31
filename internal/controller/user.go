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
	userInfo, err := service.User().Get(ctx, req)
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

func (c *cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (*v1.UserUpdateRes, error) {
	data := model.UserUpdateInput{
		Id:       req.Id,
		Nickname: req.Nickname,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	err := service.User().Update(ctx, data)

	return nil, err
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
	where := model.UserListInput{
		Page: req.Page,
		Size: req.Size,
	}
	if req.Id > 0 {
		where.Id = req.Id
	}
	if req.Nickname != "" {
		where.Nickname = req.Nickname
	}
	res, err := service.User().List(ctx, where)
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
			Total: res.Total,
			Page:  res.Page,
			Size:  res.Size,
		},
		List: list,
	}, nil
}
