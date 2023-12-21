package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoBannerOfCreate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Picture string `json:"picture" form:"picture" valid:"required,url,max=255" label:"图片"`
	Client  string `json:"client" form:"client" valid:"required,oneof=pc mobile" label:"客户端"`
	Target  string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL     string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.Order
	request.Enable
}

type DoBannerOfUpdate struct {
	Name    string `form:"name" json:"name" valid:"required,max=120" label:"名称"`
	Picture string `json:"picture" form:"picture" valid:"required,url,max=255" label:"图片"`
	Client  string `json:"client" form:"client" valid:"required,oneof=pc mobile" label:"客户端"`
	Target  string `json:"target" form:"target" valid:"required,oneof=_blank _self" label:"打开"`
	URL     string `json:"url" form:"url" valid:"omitempty,url|uri,max=255" label:"链接"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoBannerOfDelete struct {
	request.IDOfUint
}

type DoBannerOfEnable struct {
	request.IDOfUint
	request.Enable
}

type ToBannerOfPaginate struct {
	Client string `query:"client" valid:"omitempty,oneof=pc mobile" label:"客户端"`

	request.Paginate
}
