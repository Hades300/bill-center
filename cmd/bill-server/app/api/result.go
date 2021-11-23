package api

import (
	"os"
	"path"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hades300/bill-center/cmd/bill-server/app/service"
)

type ResultApiI interface {
	Parse(r *ghttp.Request)
}

type ResultApi struct{}

var _ ResultApiI = &ResultApi{}

var Result = &ResultApi{}

const uploadFileDir = "./public/resource/upload"

// get upload file
func (ra *ResultApi) Parse(r *ghttp.Request) {
	// get file and defer remove func
	file := r.GetUploadFile("file")
	filename, err := file.Save(uploadFileDir, true)
	p := path.Join(uploadFileDir, filename)
	rmFunc := func() {
		err := os.Remove(p)
		if err != nil {
			panic(err)
		}
	}
	defer rmFunc()

	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	ret, err := service.Result.Parse(p)
	if err != nil {
		JsonErrExit(r, 1, err.Error())
	}
	JsonSuccessExit(r, "", ret)
}
