package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net"
	"strings"
)

var (
	middleLogger = g.Log("Middleware")
)

func LocalhostAuth(r *ghttp.Request) {
	ipStr := r.GetRemoteIp()
	ip := net.ParseIP(strings.Trim(ipStr, "[]")) // fix ipv6 address
	if ip.IsLoopback() {
		middleLogger.Debugf(nil, "localhost user login %s", r.RemoteAddr)
		err := r.Session.Set("userId", 1)
		if err != nil {
			middleLogger.Error(nil, err)
		}
	}
	r.Middleware.Next()
}

func UserAuth(r *ghttp.Request) {
	userId, err := r.Session.Get("userId")
	if err != nil {
		middleLogger.Error(nil, err)
	}
	if userId.IsNil() {
		JsonErrExit(r, 1, "no authed user")
	} else {
		middleLogger.Debugf(nil, "user %d from %s authed", userId.Int(), r.GetRemoteIp())
	}
	r.Middleware.Next()
}

func InitSession(r *ghttp.Request) {
	_, _ = r.Session.Id()
	r.Middleware.Next()
}
