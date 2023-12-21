package web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tizips/noos.tokyo/admin/biz/common"
)

func ToSetting(c context.Context, ctx *app.RequestContext) {

	common.ToSetting(c, ctx, "web")
}

func DoSetting(c context.Context, ctx *app.RequestContext) {

	common.DoSetting(c, ctx, "web")
}
