package v1

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" v:"min:0#分页号码错误"  dc:"分页号码, 默认1"`
	Size int `json:"size" in:"query" v:"max:500#每页数量最大500条" dc:"分页数量, 最大500"`
}

type CommonPaginationRes struct {
	Total int `json:"total" dc:"总数"`
	Page  int `json:"page,omitempty" dc:"分页号码"`
	Size  int `json:"size,omitempty" dc:"分页数量"`
}
