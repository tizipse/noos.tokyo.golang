package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router(router *server.Hertz) {

	//BasicRouter(router)

	WebRouter(router)

}
