package api

import (
	"errors"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"os"
	"path"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hades300/bill-center/cmd/bill-server/app/service"
)

type ResultApiI interface {
	Parse(r *ghttp.Request)
}

type ResultApi struct{}

var _ ResultApiI = &ResultApi{}

var Result = &ResultApi{}

var (
	ErrNotAuthorized   = errors.New("无权查看")
	MaxFileTimeBufffer = time.Minute * 3
)

const uploadFileDir = "./public/resource/upload"

func (ra *ResultApi) Parse(r *ghttp.Request) {
	// get user id
	userId, err := r.Session.Get("userId")
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	var args model.ResultParseApiArgs
	err = r.ParseForm(&args)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	// verify if user own collection(add invoice to collection)
	if !service.Collection.OwnByUser(args.CollectionId, userId.String()) {
		JsonErrExit(r, 1, ErrNotAuthorized.Error())
	}
	// get upload file
	file := r.GetUploadFile("file")
	filename, err := file.Save(uploadFileDir, true)
	p := path.Join(uploadFileDir, filename)
	// remove file func
	rmFunc := func() {
		time.Sleep(MaxFileTimeBufffer)
		err := os.Remove(p)
		if err != nil {
			panic(err)
		}
	}
	go rmFunc()

	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	ret, err := service.Result.Parse(p, args.CollectionId)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	vo, err := model.ToResultVO(ret)
	JsonSuccessExit(r, "", vo)
}
