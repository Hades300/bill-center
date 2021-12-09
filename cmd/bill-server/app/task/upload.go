package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hades300/bill-center/cmd/bill-server/library/utils"
	"net/http"
	"net/url"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	accessKey = g.Cfg().MustGet(gctx.New(), "qiniu.ak").String()
	secretKey = g.Cfg().MustGet(gctx.New(), "qiniu.sk").String()
	bucket    = g.Cfg().MustGet(gctx.New(), "qiniu.bucket").String()
	domain    = g.Cfg().MustGet(gctx.New(), "qiniu.domain").String()

	formUploader = NewFormUploader()
	bucketManger = NewBucketManager()

	log = g.Log("upload-task")
)

func NewFormUploader() *storage.FormUploader {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	//构建代理client对象
	client := http.Client{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploaderEx(&cfg, &storage.Client{Client: &client})
	return formUploader
}

func NewBucketManager() *storage.BucketManager {
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager
}

func Upload(localFile string) (url string, err error) {
	hash, err := utils.CalcFileHash(localFile)
	if err != nil {
		return "", err
	}
	key := hash + gfile.Ext(localFile)
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
	}

	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
		//putExtra.NoCrc32Check = true
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		log.Errorf(gctx.New(), "Upload file:%s failed:%s", localFile, err)
		return "", err
	}
	return FormatUrl(key), nil
}

func DeleteAfterDay(u string, days int) error {
	var (
		key string
		err error
	)
	key, err = getKeyFromUrl(u)
	if err != nil {
		log.Errorf(gctx.New(), "DeleteAfterDay key:%s days:%d failed:%s", key, days, err)
		return err
	}
	return bucketManger.DeleteAfterDays(bucket, key, days)
}

func FormatUrl(key string) string {
	return fmt.Sprintf("http://%s/%s", domain, key)
}

func getKeyFromUrl(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return gstr.Trim(u.Path, "/"), nil
}
