package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"github.com/hades300/bill-center/cmd/bill-server/app/service"
)

type CollectionAPI struct{}

var Collection *CollectionAPI

func (c *CollectionAPI) Register(r *ghttp.Request) {
	var (
		args model.CreateCollectionApiArgs
		err  error
	)
	err = r.ParseForm(&args)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	code := grand.Digits(4)
	collection, err := service.Collection.RegisterCollection(args.Title, code, "1", r.GetRemoteIp(), args.Ttl)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	JsonSuccessExit(r, "成功", collection)
}

// TODO: 暂时无用户概念，根据code进行查询
func (c *CollectionAPI) Fetch(r *ghttp.Request) {
	var (
		args model.FetchCollectionApiArgs
		err  error
	)
	err = r.ParseForm(&args)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	// TODO:接口限流 code校验放内存缓存
	collection, err := service.Collection.FetchCollectionByCode(args.Code, r.RemoteAddr)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	JsonSuccessExit(r, "成功", collection)
}
