package router

import (
	"github.com/hades300/bill-center/cmd/bill-server/app/api"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	err := s.SetConfigWithMap(g.Map{
		"SessionMaxAge": time.Hour * 24,
	})
	if err != nil {
		panic(err)
	}
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
	})
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.ALL("/login", api.User.Login)
		group.ALL("/register", api.User.Register)
	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(api.InitSession, MiddlewareCORS)
		group.Middleware(api.LocalhostAuth, api.UserAuth)
		group.POST("/parse", api.Result.Parse)
	})

	s.Group("/collection", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.POST("/fetch", api.Collection.Fetch)
		group.POST("/register", api.Collection.Register)
	})
}
