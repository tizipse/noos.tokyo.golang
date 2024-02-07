package web

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
	req "github.com/tizips/noos.tokyo/web/http/request/web"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func ToPage(c context.Context, ctx *app.RequestContext) {

	var request req.ToPage

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var page model.WebPage

	fp := facades.Gorm.First(&page, "`code`=?", request.Code)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "NOT Found")
		return
	} else if fp.Error != nil {
		http.Fail(ctx, "ERROR")
		return
	}

	responses := res.ToPage{
		ID:      page.ID,
		Code:    page.Code,
		Name:    page.Name,
		Content: page.Content,
	}

	http.Success(ctx, responses)
}
