package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoTitleOfCreate struct {
	Name string `form:"name" json:"name" valid:"required,max=120" label:"名称"`

	request.Order
	request.Enable
}

type DoTitleOfUpdate struct {
	Name string `form:"name" json:"name" valid:"required,max=120" label:"名称"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoTitleOfDelete struct {
	request.IDOfUint
}

type DoTitleOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToTitleOfPaginate struct {
	request.Paginate
}
