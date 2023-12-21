package basic

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	authConstants "github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/noos.tokyo/admin/constants"
	"github.com/tizips/noos.tokyo/admin/helper/authorize"
	req "github.com/tizips/noos.tokyo/admin/http/request/basic"
	res "github.com/tizips/noos.tokyo/admin/http/response/basic"
	"github.com/tizips/noos.tokyo/model"
)

func ToAccountOfInformation(c context.Context, ctx *app.RequestContext) {

	user := authorize.User(c, ctx)

	if user.ID == "" {
		http.Unauthorized(ctx)
		return
	}

	responses := res.ToAccountOfInformation{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	if user.Username != nil {
		responses.Username = *user.Username
	}

	if user.Mobile != nil {
		responses.Mobile = *user.Mobile
	}

	if user.Email != nil {
		responses.Username = *user.Email
	}

	http.Success(ctx, responses)
}

func ToAccountOfModules(c context.Context, ctx *app.RequestContext) {

	responses := make([]res.ToAccountOfModules, 0)

	var permissions []authConstants.PermissionOfTrees

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(authorize.ID(ctx)), auth.NameOfRoleWithDeveloper()); ok {

		permissions = auth.HandlerPermissionsByTrees(constants.Permissions, 0, nil, nil, true)

	} else {

		var bindings []model.SysRoleBindPermission

		facades.Gorm.
			Where("exists (?)", facades.Gorm.
				Select("1").
				Model(&model.SysUserBindRole{}).
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), authorize.ID(ctx)).
				Where("exists (?)", facades.Gorm.
					Select("1").
					Model(&model.SysUser{}).
					Where(fmt.Sprintf("%s.user_id=%s.id  and `is_enable`=?", model.TableSysUserBindRole, model.TableSysUser), util.EnableOfYes),
				),
			).
			Find(&bindings)

		codes := make([]string, len(bindings))

		for index, item := range bindings {
			codes[index] = item.Permission
		}

		permissions = auth.HandlerPermissionsByTrees(constants.Permissions, 0, nil, codes, false)
	}

	for _, item := range permissions {
		responses = append(responses, res.ToAccountOfModules{
			Code: item.Code,
			Name: item.Name,
		})
	}

	http.Success(ctx, responses)
}

func ToAccountOfPermissions(c context.Context, ctx *app.RequestContext) {

	var request req.ToAccountOfPermissions

	if err := ctx.BindAndValidate(&request); err != nil {
		http.BadRequest(ctx, err)
		return
	}

	var permissions []authConstants.Permission

	for _, item := range constants.Permissions {
		if item.Code == request.Module {
			permissions = item.Children
		}
	}

	responses := make([]string, 0)

	var results []authConstants.PermissionsOfSimple

	if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(authorize.ID(ctx)), auth.NameOfRoleWithDeveloper()); ok {

		results = auth.HandlerPermissions(permissions, 0, []string{request.Module}, nil, true)

	} else {

		var bindings []model.SysRoleBindPermission

		facades.Gorm.
			Where("exists (?)", facades.Gorm.
				Select("1").
				Model(&model.SysUserBindRole{}).
				Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`user_id`=?", model.TableSysRoleBindPermission, model.TableSysUserBindRole, model.TableSysUserBindRole), authorize.ID(ctx)).
				Where("exists (?)", facades.Gorm.
					Select("1").
					Model(&model.SysUser{}).
					Where(fmt.Sprintf("%s.user_id=%s.id  and `is_enable`=?", model.TableSysUserBindRole, model.TableSysUser), util.EnableOfYes),
				),
			).
			Find(&bindings)

		codes := make([]string, len(bindings))

		for index, item := range bindings {
			codes[index] = item.Permission
		}

		results = auth.HandlerPermissions(permissions, 0, []string{request.Module}, codes, false)
	}

	for _, item := range results {
		responses = append(responses, item.Code)
	}

	http.Success(ctx, responses)
}
