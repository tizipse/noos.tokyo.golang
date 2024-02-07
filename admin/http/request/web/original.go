package web

import "github.com/herhe-com/framework/contracts/http/request"

type DoOriginalOfCreate struct {
	Name        string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,url" label:"头像"`
	INS         string `json:"ins" form:"ins" valid:"omitempty,max=255" label:"INS"`
	Summary     string `json:"summary" form:"summary" valid:"required,max=255" label:"简介"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"内容"`

	request.Order
	request.Enable
}

type DoOriginalOfUpdate struct {
	Name        string `json:"name" form:"name" valid:"required,max=120" label:"名称"`
	Thumb       string `json:"thumb" form:"thumb" valid:"required,max=255,url" label:"头像"`
	INS         string `json:"ins" form:"ins" valid:"omitempty,max=255" label:"INS"`
	Summary     string `json:"summary" form:"summary" valid:"required,max=255" label:"简介"`
	Title       string `json:"title" form:"title" valid:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" valid:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" valid:"omitempty,max=255" label:"描述"`
	Content     string `json:"content" form:"content" valid:"required" label:"内容"`

	request.IDOfSnowflake
	request.Order
	request.Enable
}

type DoOriginalOfEnable struct {
	request.IDOfSnowflake
	request.Enable
}

type DoOriginalOfDelete struct {
	request.IDOfSnowflake
}

type ToOriginalOfInformation struct {
	request.IDOfSnowflake
}

type ToOriginalOfPaginate struct {
	request.Paginate
}
