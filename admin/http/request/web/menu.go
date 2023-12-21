package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoMenuOfCreate struct {
	Name  string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Price string `json:"price" form:"price" valid:"required,max=16" label:"价格"`
	Type  string `json:"type" form:"type" valid:"required,oneof=cut styling spa treatment hair_color perm straight_perm" label:"类型"`

	request.Order
	request.Enable
}

type DoMenuOfUpdate struct {
	Name  string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Price string `json:"price" form:"price" valid:"required,max=16" label:"价格"`
	Type  string `json:"type" form:"type" valid:"required,oneof=cut styling spa treatment hair_color perm straight_perm" label:"类型"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoMenuOfDelete struct {
	request.IDOfUint
}

type DoMenuOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToMenuOfPaginate struct {
	Type string `query:"type" valid:"omitempty,oneof=cut styling spa treatment hair_color perm straight_perm" label:"类型"`

	request.Paginate
}
