package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/admin/helper/authorize"
	req "github.com/tizips/noos.tokyo/admin/http/request/basic"
	res "github.com/tizips/noos.tokyo/admin/http/response/basic"
	"github.com/tizips/noos.tokyo/model"
	"gorm.io/gorm"
)

func DoLoginOfAccount(c context.Context, ctx *app.RequestContext) {

	var request req.DoLoginOfAccount

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var user model.SysUser

	fu := facades.Gorm.First(&user, "`username`=? and `is_enable`=?", request.Username, util.EnableOfYes)
	if fu.Error != nil {
		http.Fail(ctx, "用户名或密码错误")
		return
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		http.Fail(ctx, "用户名或密码错误")
		return
	}

	lifetime := facades.Cfg.GetInt("jwt.lifetime")

	token, err := authorize.MakeJWT(user)
	if err != nil {
		http.Login(ctx)
		return
	}

	// 查询最高授权 平台 > 集团 > 商户 > 单店
	var bind model.SysUserBindRole

	fb := facades.Gorm.Order("`role_id` asc").First(&bind, "`user_id`=?", user.ID)
	if errors.Is(fb.Error, gorm.ErrRecordNotFound) {
		http.NotFound(ctx, "未查询到被授权的角色")
		return
	} else if fb.Error != nil {
		http.Fail(ctx, "登陆失败：%v", fb.Error)
		return
	}

	responses := res.DoLogin{
		Token:    token,
		Lifetime: lifetime,
	}

	http.Success(ctx, responses)
}

func DoLoginOfOut(c context.Context, ctx *app.RequestContext) {

	if !authorize.BlacklistOfJwt(c, ctx) {
		http.Fail(ctx, "退出失败，请稍后重试")
		return
	}

	http.Success[any](ctx)
}
