package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/noos.tokyo/admin/biz/web"
	"github.com/tizips/noos.tokyo/admin/middleware"
)

func WebRouter(router *server.Hertz) {

	route := router.Group("web")
	route.Use(middleware.Auth())
	{

		setting := route.Group("setting")
		{
			setting.GET("", middleware.Permission("web.setting.list"), web.ToSetting)
			setting.PUT("", middleware.Permission("web.setting.update"), web.DoSetting)
		}

		members := route.Group("members")
		{
			members.GET(":id", web.ToMemberOfInformation)
			members.GET("", middleware.Permission("web.member.paginate"), web.ToMemberOfPaginate)
			members.PUT(":id", middleware.Permission("web.member.update"), web.DoMemberOfUpdate)
			members.DELETE(":id", middleware.Permission("web.member.delete"), web.DoMemberOfDelete)
		}

		member := route.Group("member")
		{
			member.POST("", middleware.Permission("web.member.create"), web.DoMemberOfCreate)
			member.PUT("enable", middleware.Permission("web.member.enable"), web.DoMemberOfEnable)
		}

		originals := route.Group("originals")
		{
			originals.GET(":id", web.ToOriginalOfInformation)
			originals.GET("", middleware.Permission("web.original.paginate"), web.ToOriginalOfPaginate)
			originals.PUT(":id", middleware.Permission("web.original.update"), web.DoOriginalOfUpdate)
			originals.DELETE(":id", middleware.Permission("web.original.delete"), web.DoOriginalOfDelete)
		}

		original := route.Group("original")
		{
			original.POST("", middleware.Permission("web.original.create"), web.DoOriginalOfCreate)
			original.PUT("enable", middleware.Permission("web.original.enable"), web.DoOriginalOfEnable)
		}

		banners := route.Group("banners")
		{
			banners.GET("", middleware.Permission("web.banner.paginate"), web.ToBannerOfPaginate)
			banners.PUT(":id", middleware.Permission("web.banner.update"), web.DoBannerOfUpdate)
			banners.DELETE(":id", middleware.Permission("web.banner.delete"), web.DoBannerOfDelete)
		}

		banner := route.Group("banner")
		{
			banner.POST("", middleware.Permission("web.banner.create"), web.DoBannerOfCreate)
			banner.PUT("enable", middleware.Permission("web.banner.enable"), web.DoBannerOfEnable)
		}

		links := route.Group("links")
		{
			links.GET("", middleware.Permission("web.link.paginate"), web.ToLinkOfPaginate)
			links.PUT(":id", middleware.Permission("web.link.update"), web.DoLinkOfUpdate)
			links.DELETE(":id", middleware.Permission("web.link.delete"), web.DoLinkOfDelete)
		}

		link := route.Group("link")
		{
			link.POST("", middleware.Permission("web.link.create"), web.DoLinkOfCreate)
			link.PUT("enable", middleware.Permission("web.link.enable"), web.DoLinkOfEnable)
		}

		recruits := route.Group("recruits")
		{
			recruits.GET("", middleware.Permission("web.recruit.paginate"), web.ToRecruitOfPaginate)
			recruits.PUT(":id", middleware.Permission("web.recruit.update"), web.DoRecruitOfUpdate)
			recruits.DELETE(":id", middleware.Permission("web.recruit.delete"), web.DoRecruitOfDelete)
		}

		recruit := route.Group("recruit")
		{
			recruit.POST("", middleware.Permission("web.recruit.create"), web.DoRecruitOfCreate)
			recruit.PUT("enable", middleware.Permission("web.recruit.enable"), web.DoRecruitOfEnable)
		}

		menus := route.Group("menus")
		{
			menus.GET("", middleware.Permission("web.menu.paginate"), web.ToMenuOfPaginate)
			menus.PUT(":id", middleware.Permission("web.menu.update"), web.DoMenuOfUpdate)
			menus.DELETE(":id", middleware.Permission("web.menu.delete"), web.DoMenuOfDelete)
		}

		menu := route.Group("menu")
		{
			menu.POST("", middleware.Permission("web.menu.create"), web.DoMenuOfCreate)
			menu.PUT("enable", middleware.Permission("web.menu.enable"), web.DoMenuOfEnable)
		}

		pages := route.Group("pages")
		{
			pages.GET("", middleware.Permission("web.page.paginate"), web.ToPageOfPaginate)
			pages.PUT(":id", middleware.Permission("web.page.update"), web.DoPageOfUpdate)
			pages.DELETE(":id", middleware.Permission("web.page.delete"), web.DoPageOfDelete)
			pages.GET(":id", web.ToPageOfInformation)
		}

		page := route.Group("page")
		{
			page.POST("", middleware.Permission("web.page.create"), web.DoPageOfCreate)
		}

		titles := route.Group("titles")
		{
			titles.GET("", middleware.Permission("web.title.paginate"), web.ToTitleOfPaginate)
			titles.PUT(":id", middleware.Permission("web.title.update"), web.DoTitleOfUpdate)
			titles.DELETE(":id", middleware.Permission("web.title.delete"), web.DoTitleOfDelete)
		}

		title := route.Group("title")
		{
			title.POST("", middleware.Permission("web.title.create"), web.DoTitleOfCreate)
			title.PUT("enable", middleware.Permission("web.title.enable"), web.DoTitleOfEnable)
			title.GET("opening", web.ToTitleOfOpening)
		}

		times := route.Group("times")
		{
			times.GET("", middleware.Permission("web.time.paginate"), web.ToTimeOfPaginate)
			times.PUT(":id", middleware.Permission("web.time.update"), web.DoTimeOfUpdate)
			times.DELETE(":id", middleware.Permission("web.time.delete"), web.DoTimeOfDelete)
		}

		time := route.Group("time")
		{
			time.POST("", middleware.Permission("web.time.create"), web.DoTimeOfCreate)
			time.PUT("enable", middleware.Permission("web.time.enable"), web.DoTimeOfEnable)
		}

	}
}
