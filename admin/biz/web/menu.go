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

func DoMenuOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMenuOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	menu := model.WebMenu{
		Name:     request.Name,
		Price:    request.Price,
		Type:     request.Type,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&menu); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoMenuOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMenuOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var menu model.WebMenu

	fm := facades.Gorm.First(&menu, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	menu.Name = request.Name
	menu.Price = request.Price
	menu.Type = request.Type
	menu.Order = request.Order.Order
	menu.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&menu); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoMenuOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoMenuOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var menu model.WebMenu

	fm := facades.Gorm.First(&menu, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&menu); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoMenuOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoMenuOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var menu model.WebMenu

	fm := facades.Gorm.First(&menu, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&menu).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToMenuOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToMenuOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToMenuOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	if request.Type != "" {
		tx = tx.Where("`type` = ?", request.Type)
	}

	tx.Model(&model.WebMenu{}).Count(&responses.Total)

	if responses.Total > 0 {

		var menus []model.WebMenu

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&menus)

		responses.Data = make([]res.ToMenuOfPaginate, len(menus))

		for idx, item := range menus {
			responses.Data[idx] = res.ToMenuOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Price:     item.Price,
				Type:      item.Type,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
