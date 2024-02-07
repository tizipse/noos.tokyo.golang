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

func DoRecruitOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoRecruitOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	recruit := model.WebRecruit{
		Name:     request.Name,
		Summary:  request.Summary,
		URL:      request.URL,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := facades.Gorm.Create(&recruit); result.Error != nil {
		http.Fail(ctx, "创建失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoRecruitOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoRecruitOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var recruit model.WebRecruit

	fm := facades.Gorm.First(&recruit, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	recruit.Name = request.Name
	recruit.Summary = request.Summary
	recruit.URL = request.URL
	recruit.Order = request.Order.Order
	recruit.IsEnable = request.IsEnable

	if result := facades.Gorm.Save(&recruit); result.Error != nil {
		http.Fail(ctx, "修改失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoRecruitOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoRecruitOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var recruit model.WebRecruit

	fm := facades.Gorm.First(&recruit, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "数据查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Delete(&recruit); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func DoRecruitOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoRecruitOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var recruit model.WebRecruit

	fm := facades.Gorm.First(&recruit, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "数据不存在")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查找失败：%v", fm.Error)
		return
	}

	if result := facades.Gorm.Model(&recruit).Update("is_enable", request.IsEnable); result.Error != nil {
		http.Fail(ctx, "启禁失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToRecruitOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToRecruitOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToRecruitOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(&model.WebRecruit{}).Count(&responses.Total)

	if responses.Total > 0 {

		var recruits []model.WebRecruit

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&recruits)

		responses.Data = make([]res.ToRecruitOfPaginate, len(recruits))

		for idx, item := range recruits {
			responses.Data[idx] = res.ToRecruitOfPaginate{
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
