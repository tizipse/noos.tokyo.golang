package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
	req "github.com/tizips/noos.tokyo/web/http/request/web"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
	"gorm.io/gorm"
)

func ToOriginalOfOpening(c context.Context, ctx *app.RequestContext) {

	var originals []model.WebOriginal

	facades.Gorm.
		Order("`order` asc, `created_at` asc").
		Find(&originals, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToOriginalOfOpening, len(originals))

	for idx, item := range originals {

		responses[idx] = res.ToOriginalOfOpening{
			ID:      item.ID,
			Name:    item.Name,
			Thumb:   item.Thumb,
			Summary: item.Summary,
		}
	}

	http.Success(ctx, responses)
}

func ToOriginalOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToOriginalOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequestWithoutTranslate(ctx, err)
		return
	}

	var original model.WebOriginal

	fm := facades.Gorm.
		Preload("HTML").
		First(&original, "`id`=? and `is_enable`=?", request.ID, util.EnableOfYes)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Not Found")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "The data query failed")
		return
	}

	responses := res.ToOriginalOfInformation{
		ID:      original.ID,
		Name:    original.Name,
		Thumb:   original.Thumb,
		INS:     original.INS,
		Summary: original.Summary,
	}

	if original.HTML != nil {
		responses.Introduce = original.HTML.Content
	}

	http.Success(ctx, responses)
}
