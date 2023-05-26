package user

import (
	"context"
	"go-base/internal/dao"
	"go-base/internal/model"
	"go-base/internal/model/do"
	"go-base/internal/model/entity"
	"go-base/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Get(ctx context.Context, id uint) (*entity.User, error) {
	userInfo, err := dao.User.Ctx(ctx).Where(do.User{
		Id: id,
	}).One()
	if err != nil {
		return nil, err
	}

	var user *entity.User

	if err := userInfo.Struct(&user); err != nil {
		return nil, err
	}

	return user, err
}

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (*model.UserCreateOutput, error) {

	id, err := dao.User.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.UserCreateOutput{
		ID: uint(id),
	}, nil
}

func (s *sUser) List(ctx context.Context, in model.UserListInput) (*model.UserListOutput, error) {
	m := dao.User.Ctx(ctx)
	out := &model.UserListOutput{
		PageSize: in.PageSize,
		PageNum:  in.PageNum,
	}
	listModel := m.Page(in.PageNum, in.PageSize)

	listModel = listModel.OrderDesc(dao.User.Columns().UpdatedAt)

	if err := listModel.Scan(&out.List); err != nil {
		return nil, err
	}
	if len(out.List) == 0 {
		return &model.UserListOutput{}, nil
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	return out, nil
}
