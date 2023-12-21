package web

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/samber/lo"
	"github.com/tizips/noos.tokyo/model"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
)

func ToMenuOfOpening(c context.Context, ctx *app.RequestContext) {

	var menus []model.WebMenu

	facades.Gorm.Order("`order` asc, `id` asc").Find(&menus, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToMenuOfOpening, 0)

	for _, item := range menus {

		if _, _, ok := lo.FindIndexOf(responses, func(val res.ToMenuOfOpening) bool { return item.Type == val.Code }); !ok {

			responses = append(responses, res.ToMenuOfOpening{
				Code:     item.Type,
				Label:    item.ToType(),
				Children: make([]res.ToMenuOfOpeningOfChildren, 0),
			})
		}
	}

	for idx, item := range responses {

		for _, val := range menus {

			if item.Code == val.Type {
				responses[idx].Children = append(responses[idx].Children, res.ToMenuOfOpeningOfChildren{
					ID:    val.ID,
					Name:  val.Name,
					Price: val.Price,
				})
			}
		}
	}

	http.Success(ctx, responses)
}
