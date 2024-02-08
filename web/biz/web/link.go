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

func ToLinkOfOpening(c context.Context, ctx *app.RequestContext) {

	var links []model.WebLink

	facades.Gorm.Order("`order` asc, `id` asc").Find(&links, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToLinkOfOpening, len(links))

	for idx, item := range links {
		responses[idx] = res.ToLinkOfOpening{
			ID:       item.ID,
			Name:     item.Name,
			Summary:  item.Summary,
			URL:      item.URL,
			IsSystem: item.IsSystem,
		}
	}

	http.Success(ctx, responses)
}
