package model

import "github.com/gogf/gf/v2/os/gtime"

type UserCreateInput struct {
	Nickname string
	Age      int
	Gender   int
}

type UserCreateOutput struct {
	ID uint
}

type UserListInput struct {
	PageNum  int
	PageSize int
}

type UserListOutput struct {
	PageNum  int
	PageSize int
	Total    int
	List     []UserListOutputItem
}

// DemoListOutputItem NOTE: 此处为了不返回Fieldb字段, 所以重新定义返回结构体, 否则可以直接使用enttity.Demo
type UserListOutputItem struct {
	Id        uint        `json:"id"`
	Nickname  string      `json:"nickname"`
	Age       int         `json:"age"`
	Gender    int         `json:"gender"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}
