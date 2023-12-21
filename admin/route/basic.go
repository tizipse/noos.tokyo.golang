package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/noos.tokyo/admin/biz/basic"
	"github.com/tizips/noos.tokyo/admin/middleware"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{

		login := route.Group("login")
		{
			login.POST("account", basic.DoLoginOfAccount)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.GET("information", basic.ToAccountOfInformation)
			account.GET("modules", basic.ToAccountOfModules)
			account.GET("permissions", basic.ToAccountOfPermissions)
			account.POST("logout", middleware.Auth(), basic.DoLoginOfOut)
		}

		upload := route.Group("upload").Use(middleware.Auth())
		{
			upload.POST("file", basic.DoUploadOfFile)
		}
	}
}
