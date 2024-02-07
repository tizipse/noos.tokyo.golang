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

		recruit := route.Group("recruit")
		{
			recruit.GET("opening", web.ToRecruitOfOpening)
		}

		link := route.Group("link")
		{
			link.GET("opening", web.ToLinkOfOpening)
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

		originals := route.Group("originals")
		{
			originals.GET(":id", web.ToOriginalOfInformation)
		}

		original := route.Group("original")
		{
			original.GET("opening", web.ToOriginalOfOpening)
		}

		seo := route.Group("seo")
		{
			seo.GET("", web.ToSEO)
		}

		page := route.Group("page")
		{
			page.GET("", web.ToPage)
		}

		setting := route.Group("setting")
		{
			setting.GET("", web.ToSetting)
		}
	}
}
