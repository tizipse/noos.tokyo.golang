package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoTimeOfCreate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Status  string `form:"status" json:"status" valid:"required,oneof=open close" label:"状态"`
	Content string `form:"content" json:"content" valid:"required_if=Status open,max=120" label:"内容"`

	request.Order
	request.Enable
}

type DoTimeOfUpdate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Status  string `form:"status" json:"status" valid:"required,oneof=open close" label:"状态"`
	Content string `form:"content" json:"content" valid:"required_if=Status open,max=120" label:"内容"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoTimeOfDelete struct {
	request.IDOfUint
}

type DoTimeOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToTimeOfPaginate struct {
	request.Paginate
}
