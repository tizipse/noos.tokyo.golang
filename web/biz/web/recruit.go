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

func ToRecruitOfOpening(c context.Context, ctx *app.RequestContext) {

	var recruits []model.WebRecruit

	facades.Gorm.Order("`order` asc, `id` asc").Find(&recruits, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToRecruitOfOpening, len(recruits))

	for idx, item := range recruits {
		responses[idx] = res.ToRecruitOfOpening{
			ID:      item.ID,
			Name:    item.Name,
			Summary: item.Summary,
			URL:     item.URL,
		}
	}

	http.Success(ctx, responses)
}
