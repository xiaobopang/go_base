package user

import (
	"context"
	v1 "go-base/api/v1"
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

func (s *sUser) Get(ctx context.Context, req *v1.UserGetReq) (*entity.User, error) {
	userInfo, err := dao.User.Ctx(ctx).Where("id=?", req.Id).One()
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

func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) error {
	_, err := dao.User.Ctx(ctx).Fields("nickname", "age", "gender").Data(in).Where(do.Demo{
		Id: in.Id,
	}).Update()
	if err != nil {
		return err
	}

	return nil
}
func (s *sUser) List(ctx context.Context, in model.UserListInput) (*model.UserListOutput, error) {
	out := &model.UserListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := dao.User.Ctx(ctx)

	if in.Id > 0 {
		listModel = listModel.Where("id=?", in.Id)
	}

	if in.Nickname != "" {
		listModel = listModel.WhereLike("nickname", "%"+in.Nickname+"%")
	}

	if err := listModel.Page(in.Page, in.Size).OrderDesc(dao.User.Columns().UpdatedAt).Scan(&out.List); err != nil {
		return nil, err
	}

	if len(out.List) == 0 {
		return &model.UserListOutput{}, nil
	}
	count, err := listModel.Count()
	if err != nil {
		return nil, err
	}
	out.Total = count

	return out, nil
}
