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

func DoBannerOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	banner := model.WebBanner{
		Name:     request.Name,
		Picture:  request.Picture,
		Client:   request.Client,
		Target:   request.Target,
		URL:      request.URL,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&banner); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.WebBanner

	fm := facades.Gorm.First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	banner.Name = request.Name
	banner.Picture = request.Picture
	banner.Client = request.Client
	banner.Target = request.Target
	banner.URL = request.URL
	banner.Order = request.Order.Order
	banner.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&banner); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.WebBanner

	fm := facades.Gorm.First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&banner); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoBannerOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoBannerOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var banner model.WebBanner

	fm := facades.Gorm.First(&banner, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&banner).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToBannerOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToBannerOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToBannerOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	if request.Client != "" {
		tx = tx.Where("client = ?", request.Client)
	}

	tx.Model(&model.WebBanner{}).Count(&responses.Total)

	if responses.Total > 0 {

		var banners []model.WebBanner

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&banners)

		responses.Data = make([]res.ToBannerOfPaginate, len(banners))

		for idx, item := range banners {
			responses.Data[idx] = res.ToBannerOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Picture:   item.Picture,
				Client:    item.Client,
				Target:    item.Target,
				URL:       item.URL,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
