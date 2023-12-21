package web

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/model"
	req "github.com/tizips/noos.tokyo/web/http/request/web"
	res "github.com/tizips/noos.tokyo/web/http/response/web"
	"gorm.io/gorm"
)

func ToMemberOfOpening(c context.Context, ctx *app.RequestContext) {

	var members []model.WebMember

	facades.Gorm.
		Preload("Title", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Order("`order` asc, `created_at` asc").
		Find(&members, "`is_enable`=?", util.EnableOfYes)

	responses := make([]res.ToMemberOfOpening, len(members))

	for idx, item := range members {

		responses[idx] = res.ToMemberOfOpening{
			ID:         item.ID,
			Name:       item.Name,
			Thumb:      item.Thumb,
			IsDelegate: item.IsDelegate,
		}

		if item.Title != nil {
			responses[idx].Title = item.Title.Name
		}
	}

	http.Success(ctx, responses)
}

func ToMemberOfInformation(c context.Context, ctx *app.RequestContext) {

	var request req.ToMemberOfInformation

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequestWithoutTranslate(ctx, err)
		return
	}

	var member model.WebMember

	fm := facades.Gorm.
		Preload("Title", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Preload("HTML").
		First(&member, "`id`=? and `is_enable`=?", request.ID, util.EnableOfYes)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "Not Found")
		return
	} else if fm.Error != nil {
		http.Fail(ctx, "The data query failed")
		return
	}

	responses := res.ToMemberOfInformation{
		ID:         member.ID,
		Name:       member.Name,
		Nickname:   member.Nickname,
		Thumb:      member.Thumb,
		INS:        member.INS,
		IsDelegate: member.IsDelegate,
	}

	if member.Title != nil {
		responses.Title = member.Title.Name
	}

	if member.HTML != nil {
		responses.Introduce = member.HTML.Content
	}

	http.Success(ctx, responses)
}
