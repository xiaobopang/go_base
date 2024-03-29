// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-05-30 13:33:02
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        ` //
	Nickname  string      `json:"nickname"  ` //
	Age       uint        `json:"age"       ` //
	Gender    uint        `json:"gender"    ` //
	Status    uint        `json:"status"    ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
