package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/noos.tokyo/web/biz/web"
)

func WebRouter(routes *server.Hertz) {

	route := routes.Group("web")
	{

		banner := route.Group("banner")
		{
			banner.GET("opening", web.ToBannerOfOpening)
		}

		menu := route.Group("menu")
		{
			menu.GET("opening", web.ToMenuOfOpening)
		}

		time := route.Group("time")
		{
			time.GET("opening", web.ToTimeOfOpening)
		}

		members := route.Group("members")
		{
			members.GET(":id", web.ToMemberOfInformation)
		}

		member := route.Group("member")
		{
			member.GET("opening", web.ToMemberOfOpening)
		}

		seo := route.Group("seo")
		{
			seo.GET("", web.ToSEO)
		}

		setting := route.Group("setting")
		{
			setting.GET("", web.ToSetting)
		}
	}
}
