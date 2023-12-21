package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
)

func ToSetting(c context.Context, ctx *app.RequestContext, module string) {

	var settings []model.ComSetting

	facades.Gorm.Order("`order` asc, `id` asc").Find(&settings, "`module`=?", module)

	responses := make(map[string]string, len(settings))

	for _, item := range settings {
		responses[item.Key] = item.Val
	}

	http.Success(ctx, responses)
}
