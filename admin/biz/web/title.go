package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/tizips/noos.tokyo/admin/http/request/web"
	res "github.com/tizips/noos.tokyo/admin/http/response/web"
	"github.com/tizips/noos.tokyo/model"
	"gorm.io/gorm"
)

func DoTitleOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoTitleOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	title := model.WebTitle{
		Name:     request.Name,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&title); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTitleOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoTitleOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var title model.WebTitle

	ft := facades.Gorm.First(&title, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "查找失败：%v", ft.Error)
		return
	}

	title.Name = request.Name
	title.Order = request.Order.Order
	title.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&title); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTitleOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoTitleOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var title model.WebTitle

	ft := facades.Gorm.First(&title, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", ft.Error)
		return
	}

	if result := facades.Gorm.Delete(&title); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTitleOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoTitleOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var title model.WebTitle

	ft := facades.Gorm.First(&title, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "查找失败：%v", ft.Error)
		return
	}

	if result := facades.Gorm.Model(&title).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToTitleOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToTitleOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToTitleOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(&model.WebTitle{}).Count(&responses.Total)

	if responses.Total > 0 {

		var titles []model.WebTitle

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&titles)

		responses.Data = make([]res.ToTitleOfPaginate, len(titles))

		for idx, item := range titles {
			responses.Data[idx] = res.ToTitleOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}

func ToTitleOfOpening(c context.Context, ctx *app.RequestContext) {

	var titles []model.WebTitle

	facades.Gorm.Order("`order` asc, `id` asc").Find(&titles, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToTitleOfOpening, len(titles))

	for idx, item := range titles {
		responses[idx] = res.ToTitleOfOpening{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(ctx, responses)
}
