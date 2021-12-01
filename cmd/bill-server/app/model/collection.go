package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hades300/bill-center/cmd/bill-server/library/convert"
)

type CreateCollectionApiArgs struct {
	Title string `v:"required#标题是必要的`   // 集合标题
	Code  string `v:"required#分享码是必要的`  // 集合密码
	Ttl   int    `v:"required#有效期是必要的"` // 有效期（单位秒）
}

type FetchCollectionApiArgs struct {
	Code string `v:"required#分享码是必要的"` // 集合密码
}

type CollectionVO struct {
	Title       string      `json:"title"       ` // 集合标题
	UserId      int         `json:"userId"      ` // 外键 留用
	GmtCreated  *gtime.Time `json:"gmtCreated"  ` // 创建UTC时间
	GmtModified *gtime.Time `json:"gmtModified" ` // 更新UTC时间

	ResultList []*ResultVO `json:"resultList" ` // 发票列表
}

func ToCollectionVO(collection *Collection, resultList []*Result) *CollectionVO {
	var ret CollectionVO
	convert.Transform(collection, &ret)
	var resList []*ResultVO
	for _, v := range resultList {
		var tmp ResultVO
		convert.Transform(v, &tmp)
		resList = append(resList, &tmp)
	}
	ret.ResultList = resList
	return &ret
}
