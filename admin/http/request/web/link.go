package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoLinkOfCreate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Summary string `form:"summary" json:"summary" valid:"required,max=255" label:"简介"`
	URL     string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.Order
	request.Enable
}

type DoLinkOfUpdate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Summary string `form:"summary" json:"summary" valid:"required,max=255" label:"简介"`
	URL     string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoLinkOfDelete struct {
	request.IDOfUint
}

type DoLinkOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToLinkOfPaginate struct {
	request.Paginate
}
