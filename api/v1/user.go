package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go-base/internal/model"
)

// 查询记录示例
type UserGetReq struct {
	g.Meta `path:"/user/:id" method:"post" tags:"UserService" summary:"Get user info by id"`
	Id     uint `p:"id" in:"path" v:"required|integer#id不能为空"`
}

type UserGetRes struct {
	Id        uint        `json:"id"        ` //
	Nickname  string      `json:"nickname"  ` //
	Age       uint        `json:"age"       ` //
	Gender    uint        `json:"gender"    ` //
	Status    uint        `json:"status"    ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}

// 创建记录示例
type UserCreateReq struct {
	g.Meta   `path:"/user-create" method:"post" tags:"UserService" summary:"Create a user record"`
	Nickname string `p:"nickname" v:"required"`
	Age      int    `p:"age" v:"required|integer"`
	Gender   int    `p:"gender" v:"required|integer"`
}

type UserCreateRes struct {
	ID uint `json:"id"`
}

// 查询记录列表示例
type UserListReq struct {
	g.Meta   `path:"/user-list" method:"post" tags:"UserService" summary:"Get user records list"`
	Id       uint   `p:"id"`
	Nickname string `p:"nickname"`
	CommonPaginationReq
}

type UserListRes struct {
	CommonPaginationRes

	List []model.UserListOutputItem `json:"list"`
}

// 更新记录示例
type UserUpdateReq struct {
	g.Meta   `path:"/user-update/:id" method:"put" tags:"UserService" summary:"Update a user record"`
	Id       uint   `p:"id" in:"path" v:"integer|min:1"`
	Nickname string `p:"nickname"`
	Age      int    `p:"age"`
	Gender   int    `p:"gender" v:"between:0,3"`
}

type UserUpdateRes struct {
}
