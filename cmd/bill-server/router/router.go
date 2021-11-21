package router

import (
	"github.com/hades300/bill-center/cmd/bill-server/app/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
	})
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.ALL("/login", api.User.Login)
		group.ALL("/register", api.User.Register)
	})
}
