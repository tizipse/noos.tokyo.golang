package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/tizips/noos.tokyo/admin/http/request/web"
	res "github.com/tizips/noos.tokyo/admin/http/response/web"
	"github.com/tizips/noos.tokyo/model"
	"gorm.io/gorm"
)

func DoLinkOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoLinkOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	link := model.WebLink{
		Name:     request.Name,
		Summary:  request.Summary,
		URL:      request.URL,
		Order:    request.Order.Order,
		IsSystem: model.WebLinkOfIsSystemNO,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&link); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoLinkOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoLinkOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var link model.WebLink

	fm := facades.Gorm.First(&link, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	link.Name = request.Name
	link.Summary = request.Summary
	link.URL = request.URL
	link.Order = request.Order.Order
	link.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&link); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoLinkOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoLinkOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var link model.WebLink

	fm := facades.Gorm.First(&link, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&link); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoLinkOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoLinkOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var link model.WebLink

	fm := facades.Gorm.First(&link, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&link).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToLinkOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToLinkOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToLinkOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(&model.WebLink{}).Count(&responses.Total)

	if responses.Total > 0 {

		var links []model.WebLink

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&links)

		responses.Data = make([]res.ToLinkOfPaginate, len(links))

		for idx, item := range links {
			responses.Data[idx] = res.ToLinkOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Summary:   item.Summary,
				URL:       item.URL,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
