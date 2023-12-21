package bootstrap

import (

	//Containers and other services must be started immediately
	"github.com/herhe-com/framework/foundation"

	//Delayed startup of other services init
	"github.com/tizips/noos.tokyo/admin/config"
)

func Boot() {

	application := foundation.Application{}

	//Bootstrap the application.
	application.Boot()

	//Bootstrap the other service.
	config.Boot()
}
