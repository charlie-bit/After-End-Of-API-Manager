package models

type Pager struct {
	Page int  `json:"page"`//数据页码
	Total int //数据总页数
	Size int //每页数据条数
}
