package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoMemberOfCreate struct {
	TitleID     uint   `json:"title_id" form:"title_id" valid:"required,gt=0" label:"职位"`
	Name        string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Nickname    string `json:"nickname" form:"nickname" valid:"required,max=120" label:"别称"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,url" label:"头像"`
	INS         string `json:"ins" form:"ins" valid:"omitempty,max=255" label:"INS"`
	IsDelegate  uint8  `json:"is_delegate" form:"is_delegate" valid:"required,oneof=1 2" label:"是否代表"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"内容"`

	request.Order
	request.Enable
}

type DoMemberOfUpdate struct {
	TitleID     uint   `json:"title_id" form:"title_id" valid:"required,gt=0" label:"职位"`
	Name        string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Nickname    string `json:"nickname" form:"nickname" valid:"required,max=120" label:"别称"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,url" label:"头像"`
	INS         string `json:"ins" form:"ins" valid:"omitempty,max=255" label:"INS"`
	IsDelegate  uint8  `json:"is_delegate" form:"is_delegate" valid:"required,oneof=1 2" label:"是否代表"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"内容"`

	request.IDOfSnowflake
	request.Order
	request.Enable
}

type DoMemberOfEnable struct {
	request.IDOfSnowflake
	request.Enable
}

type DoMemberOfDelete struct {
	request.IDOfSnowflake
}

type ToMemberOfInformation struct {
	request.IDOfSnowflake
}

type ToMemberOfPaginate struct {
	request.Paginate
}
