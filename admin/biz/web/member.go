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

func DoMemberOfCreate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMemberOfCreate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var title model.WebTitle

	ft := facades.Gorm.First(&title, "`id`=?", request.TitleID)

	if errors.Is(ft.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该职位")
		return
	} else if ft.Error != nil {
		http.Fail(ctx, "职位查询失败：%v", ft.Error)
		return
	}

	tx := facades.Gorm.Begin()

	member := model.WebMember{
		ID:         facades.Snowflake.Generate().String(),
		TitleID:    request.TitleID,
		Name:       request.Name,
		Nickname:   request.Nickname,
		Thumb:      request.Thumb,
		INS:        request.INS,
		Order:      request.Order.Order,
		IsDelegate: request.IsDelegate,
		IsEnable:   request.IsEnable,
	}

	if result := tx.Create(&member); result.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "写入失败：%v", result.Error)
		return
	}

	seo := model.WebSEO{
		Channel:     web.ChannelOfMember,
		ChannelID:   member.ID,
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
		Channel:   web.ChannelOfMember,
		ChannelID: member.ID,
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

func DoMemberOfUpdate(c context.Context, ctx *app.RequestContext) {

	var request req.DoMemberOfUpdate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	mutex := facades.Locker.NewMutex(locker.Keys("web", "member", request.ID))

	if err := mutex.Lock(); err != nil {
		http.Fail(ctx, "处理失败：%v", err)
		return
	}

	defer func(lock *redsync.Mutex) {
		_, _ = lock.Unlock()
	}(mutex)

	var member model.WebMember

	fm := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfMember) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfMember) }).
		First(&member, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查询失败：%v", fm.Error)
		return
	}

	if request.TitleID != member.TitleID {

		var title model.WebTitle

		fc := facades.Gorm.First(&title, "`id`=?", request.TitleID)

		if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
			http.NotFound(ctx, "未找到该职位")
			return
		} else if fc.Error != nil {
			http.Fail(ctx, "职位查询失败：%v", fc.Error)
			return
		}
	}

	tx := facades.Gorm.Begin()

	member.TitleID = request.TitleID
	member.Name = request.Name
	member.Nickname = request.Nickname
	member.Thumb = request.Thumb
	member.INS = request.INS
	member.Order = request.Order.Order
	member.IsDelegate = request.IsDelegate
	member.IsEnable = request.IsEnable

	if ua := tx.Omit(clause.Associations).Save(&member); ua.Error != nil {
		tx.Rollback()
		http.Fail(ctx, "修改失败：%v", ua.Error)
		return
	}

	if member.SEO != nil {

		member.SEO.Title = request.Title
		member.SEO.Keyword = request.Keyword
		member.SEO.Description = request.Description

		if us := tx.Save(&member.SEO); us.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", us.Error)
			return
		}
	}

	if member.HTML != nil {

		member.HTML.Content = request.Content

		if ut := tx.Save(&member.HTML); ut.Error != nil {
			tx.Rollback()
			http.Fail(ctx, "修改失败：%v", ut.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](ctx)
}

func DoMemberOfEnable(c context.Context, ctx *app.RequestContext) {

	var request req.DoMemberOfEnable

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var member model.WebMember

	fm := facades.Gorm.First(&member, "`id`=?", request.ID)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "查询失败：%v", fm.Error)
		return
	}

	if ec := facades.Gorm.Model(&member).Update("is_enable", request.IsEnable); ec.Error != nil {
		http.Fail(ctx, "操作失败：%v", ec.Error)
		return
	}

	http.Success[any](ctx)
}

func DoMemberOfDelete(c context.Context, ctx *app.RequestContext) {

	var request req.DoMemberOfDelete

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var member model.WebMember

	fc := facades.Gorm.First(&member, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fc.Error != nil {
		http.Fail(ctx, "查询失败：%v", fc.Error)
		return
	}

	if result := facades.Gorm.Delete(&member); result.Error != nil {
		http.Fail(ctx, "删除失败：%v", result.Error)
		return
	}

	http.Success[any](ctx)
}

func ToMemberOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToMemberOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var member model.WebMember

	fc := facades.Gorm.
		//Preload("Title", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfMember) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`channel`=?", web.ChannelOfMember) }).
		First(&member, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未找到该数据")
		return
	} else if fc.Error != nil {
		http.Fail(ctx, "查询失败：%v", fc.Error)
		return
	}

	responses := res.ToMemberOfInformation{
		ID:         member.ID,
		TitleID:    member.TitleID,
		Name:       member.Name,
		Nickname:   member.Nickname,
		Thumb:      member.Thumb,
		INS:        member.INS,
		Order:      member.Order,
		IsDelegate: member.IsDelegate,
		IsEnable:   member.IsEnable,
	}

	if member.SEO != nil {
		responses.Title = member.SEO.Title
		responses.Keyword = member.SEO.Keyword
		responses.Description = member.SEO.Description
	}

	if member.HTML != nil {
		responses.Content = member.HTML.Content
	}

	http.Success(ctx, responses)
}

func ToMemberOfPaginate(c context.Context, ctx *app.RequestContext) {

	var request req.ToMemberOfPaginate

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	responses := response.Paginate[res.ToMemberOfPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.WithContext(c)

	tx.Model(model.WebMember{}).Count(&responses.Total)

	if responses.Total > 0 {

		var members []model.WebMember

		tx.
			Preload("Title", func(tf *gorm.DB) *gorm.DB { return tf.Unscoped() }).
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` desc").
			Find(&members)

		responses.Data = make([]res.ToMemberOfPaginate, len(members))

		for idx, item := range members {

			responses.Data[idx] = res.ToMemberOfPaginate{
				ID:         item.ID,
				Name:       item.Name,
				Nickname:   item.Nickname,
				Order:      item.Order,
				IsDelegate: item.IsDelegate,
				IsEnable:   item.IsEnable,
				CreatedAt:  item.CreatedAt.ToDateTimeString(),
			}

			if item.Title != nil {
				responses.Data[idx].Title = item.Title.Name
			}
		}
	}

	http.Success(ctx, responses)
}
