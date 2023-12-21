package config

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/herhe-com/framework/console"
	"github.com/herhe-com/framework/console/consoles"
	cons "github.com/herhe-com/framework/contracts/console"
	"github.com/herhe-com/framework/contracts/service"
	"github.com/herhe-com/framework/database/gorm"
	"github.com/herhe-com/framework/database/redis"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/herhe-com/framework/validation"
	"github.com/tizips/noos.tokyo/web/route"
	"time"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("app", map[string]any{
		"name":     cfg.Env("app.name", "UPER"),
		"address":  cfg.Env("app.address", "0.0.0.0"),
		"port":     cfg.Env("app.port", "9600"),
		"node":     cfg.Env("app.node", 1),
		"debug":    cfg.Env("app.debug", false),
		"domain":   cfg.Env("app.domain", "http://127.0.0.1:9600"),
		"location": cfg.Env("app.location", "Asia/Shanghai"),
		"providers": []service.Provider{
			&gorm.ServiceProvider{},
			&redis.ServiceProvider{},
			//&filesystem.ServiceProvider{},
			//&snowflake.ServiceProvider{},
			//&locker.ServiceProvider{},
			&validation.ServiceProvider{},
			//&queue.ServiceProvider{},
			&console.ServiceProvider{},
		},
		"consoles": []cons.Provider{
			//&consoles.MigrationProvider{},
			&consoles.ServerProvider{},
			//&consoles2.RoleProvider{},
			//&consoles2.DeveloperProvider{},
		},
		"server": map[string]any{
			"route": route.Router,
			//"handle":  func(server *server.Hertz) {},
			"options": []config.Option{
				server.WithReadTimeout(time.Minute * 3),
				server.WithWriteTimeout(time.Minute * 3),
			},
			"middlewares": []app.HandlerFunc{
				middleware.Access(),
				middleware.Cors(),
			},
		},
	})
}

func Boot() {

}
