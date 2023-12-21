package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
	req "github.com/tizips/noos.tokyo/web/http/request/web"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
	"gorm.io/gorm"
)

func ToSEO(c context.Context, ctx *app.RequestContext) {

	var request req.ToSEO

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequestWithoutTranslate(ctx, err)
		return
	}

	var seo model.WebSEO

	fs := facades.Gorm.First(&seo, "`channel`=? and `channel_id`=?", request.Channel, request.ID)

	if errors.Is(fs.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Not Found")
		return
	} else if fs.Error != nil {
		http.Fail(ctx, "The data query failed")
		return
	}

	responses := res.ToSEO{
		Title:       seo.Title,
		Keyword:     seo.Keyword,
		Description: seo.Description,
	}

	http.Success(ctx, responses)
}
