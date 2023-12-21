package web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
)

func ToTimeOfOpening(c context.Context, ctx *app.RequestContext) {

	var times []model.WebTime

	facades.Gorm.Order("`order` asc, `id` asc").Find(&times, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToTimeOfOpening, len(times))

	for idx, item := range times {
		responses[idx] = res.ToTimeOfOpening{
			ID:      item.ID,
			Name:    item.Name,
			Content: item.Content,
			Status:  item.Status,
		}
	}

	http.Success(ctx, responses)
}
