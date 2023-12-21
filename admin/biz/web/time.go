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

func DoTimeOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoTimeOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	time := model.WebTime{
		Name:     request.Name,
		Content:  request.Content,
		Status:   request.Status,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&time); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTimeOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoTimeOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var time model.WebTime

	ft := facades.Gorm.First(&time, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "查找失败：%v", ft.Error)
		return
	}

	time.Name = request.Name
	time.Content = request.Content
	time.Status = request.Status
	time.Order = request.Order.Order
	time.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&time); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTimeOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoTimeOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var time model.WebTime

	ft := facades.Gorm.First(&time, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", ft.Error)
		return
	}

	if result := facades.Gorm.Delete(&time); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoTimeOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoTimeOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var time model.WebTime

	ft := facades.Gorm.First(&time, "`id`=?", request.ID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "查找失败：%v", ft.Error)
		return
	}

	if result := facades.Gorm.Model(&time).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToTimeOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToTimeOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToTimeOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(&model.WebTime{}).Count(&responses.Total)

	if responses.Total > 0 {

		var times []model.WebTime

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&times)

		responses.Data = make([]res.ToTimeOfPaginate, len(times))

		for idx, item := range times {
			responses.Data[idx] = res.ToTimeOfPaginate{
				ID:        item.ID,
				Name:      item.Name,
				Content:   item.Content,
				Status:    item.Status,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(ctx, responses)
}
