package api

import "github.com/gogf/gf/v2/net/ghttp"

type H = map[string]interface{}

func Json(r *ghttp.Request, code int, msg string, data interface{}) {
	ret := H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	r.Response.WriteJsonExit(ret)
}

func JsonErrExit(r *ghttp.Request, code int, msg string) {
	Json(r, code, msg, nil)
}

func JsonSuccessExit(r *ghttp.Request, msg string, data interface{}) {
	Json(r, 0, msg, data)
}
