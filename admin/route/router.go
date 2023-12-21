package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/noos.tokyo/admin/middleware"
)

func Router(router *server.Hertz) {

	router.Use(middleware.Jwt())

	BasicRouter(router)

	SiteRouter(router)

	WebRouter(router)

}
