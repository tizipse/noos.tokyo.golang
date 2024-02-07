package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redsync/redsync/v4"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	req "github.com/tizips/noos.tokyo/admin/http/request/web"
	res "github.com/tizips/noos.tokyo/admin/http/response/web"
	"github.com/tizips/noos.tokyo/constants/web"
	"github.com/tizips/noos.tokyo/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DoOriginalOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoOriginalOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	tx := facades.Gorm.Begin()

	original := model.WebOriginal{
		ID:       facades.Snowflake.Generate().String(),
		Name:     request.Name,
		Thumb:    request.Thumb,
		INS:      request.INS,
		Summary:  request.Summary,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if result := tx.Create(&original); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "写入失败：%v", result.Error)
		return
	}

	seo := model.WebSEO{
		Channel:     web.ChannelOfOriginal,
		ChannelID:   original.ID,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
	}

	if result := tx.Create(&seo); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "写入失败：%v", result.Error)
		return
	}

	html := model.WebHTML{
		Channel:   web.ChannelOfOriginal,
		ChannelID: original.ID,
		Content:   request.Content,
	}

	if result := tx.Create(&html); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "写入失败：%v", result.Error)
		return
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoOriginalOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoOriginalOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	mutex := facades.Locker.NewMutex(locker.Keys("web", "original", request.ID))

	if err := mutex.Lock(); err != nil {
		http.Fail(ctx, "处理失败：%v", err)
		return
	}

	defer func(lock *redsync.Mutex) {
		_, _ = lock.Unlock()
	}(mutex)

	var original model.WebOriginal

	fm := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfOriginal) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfOriginal) }).
		First(&original, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查询失败：%v", fm.Error)
		return
	}

	tx := facades.Gorm.Begin()

	original.Name = request.Name
	original.Thumb = request.Thumb
	original.INS = request.INS
	original.Summary = request.Summary
	original.Order = request.Order.Order
	original.IsEnable = request.IsEnable

	if ua := tx.Omit(clause.Associations).Save(&original); ua.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", ua.Error)
		return
	}

	if original.SEO != nil {

		original.SEO.Title = request.Title
		original.SEO.Keyword = request.Keyword
		original.SEO.Description = request.Description

		if us := tx.Save(&original.SEO); us.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", us.Error)
			return
		}
	}

	if original.HTML != nil {

		original.HTML.Content = request.Content

		if ut := tx.Save(&original.HTML); ut.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", ut.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoOriginalOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoOriginalOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var original model.WebOriginal

	fm := facades.Gorm.First(&original, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查询失败：%v", fm.Error)
		return
	}

	if ec := facades.Gorm.Model(&original).Update("is_enable", request.IsEnable); ec.Error != nil {
		http.Fail(ctx, "操作失败：%v", ec.Error)
		return
	}

	http.Success[any](ctx)
}

func DoOriginalOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoOriginalOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var original model.WebOriginal

	fc := facades.Gorm.First(&original, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fc.Error != nil {
		http.Fail(ctx, "查询失败：%v", fc.Error)
		return
	}

	if result := facades.Gorm.Delete(&original); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToOriginalOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToOriginalOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var original model.WebOriginal

	fc := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfOriginal) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfOriginal) }).
		First(&original, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fc.Error != nil {
		http.Fail(ctx, "查询失败：%v", fc.Error)
		return
	}

	responses := res.ToOriginalOfInformation{
		ID:       original.ID,
		Name:     original.Name,
		Thumb:    original.Thumb,
		INS:      original.INS,
		Summary:  original.Summary,
		Order:    original.Order,
		IsEnable: original.IsEnable,
	}

	if original.SEO != nil {
		responses.Title = original.SEO.Title
		responses.Keyword = original.SEO.Keyword
		responses.Description = original.SEO.Description
	}

	if original.HTML != nil {
		responses.Content = original.HTML.Content
	}

	http.Success(ctx, responses)
}

func ToOriginalOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToOriginalOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToOriginalOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(model.WebOriginal{}).Count(&responses.Total)

	if responses.Total > 0 {

		var originals []model.WebOriginal

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` desc").
			Find(&originals)

		responses.Data = make([]res.ToOriginalOfPaginate, len(originals))

		for idx, item := range originals {

			responses.Data[idx] = res.ToOriginalOfPaginate{
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
