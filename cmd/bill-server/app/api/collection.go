package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"github.com/hades300/bill-center/cmd/bill-server/app/service"
)

type CollectionAPI struct{}

var Collection *CollectionAPI

func (c *CollectionAPI) register(r *ghttp.Request) {
	var (
		args model.CreateCollectionApiArgs
		err  error
	)
	err = r.ParseForm(&args)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	err = service.Collection.RegisterCollection(args.Code, "1", r.RemoteAddr, args.Ttl)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	JsonSuccessExit(r, "成功", nil)
}

// TODO: 暂时无用户概念，根据code进行查询
func (c *CollectionAPI) fetch(r *ghttp.Request) {
	var (
		args model.FetchCollectionApiArgs
		err  error
	)
	err = r.ParseForm(&args)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	collection, err := service.Collection.FetchCollectionByCode(args.Code, r.RemoteAddr)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	JsonSuccessExit(r, "成功", collection)
}
