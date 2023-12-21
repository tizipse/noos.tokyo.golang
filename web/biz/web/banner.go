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

func ToBannerOfOpening(c context.Context, ctx *app.RequestContext) {

	var banners []model.WebBanner

	facades.Gorm.Order("`order` asc, `id` asc").Find(&banners, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToBannerOfOpening, len(banners))

	for idx, item := range banners {
		responses[idx] = res.ToBannerOfOpening{
			ID:      item.ID,
			Client:  item.Client,
			Name:    item.Name,
			Picture: item.Picture,
			Target:  item.Target,
			URL:     item.URL,
		}
	}

	http.Success(ctx, responses)
}
