package service

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hades300/bill-center/cmd/bill-server/library/convert"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hades300/bill-center/cmd/bill-server/app/dao"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
)

type collectionServiceI interface {
	FetchCollectionByCode(code string, userIp string) (*model.CollectionVO, error)
	RegisterCollection(title string, code string, userId string, userIp string, ttl int) (*model.Collection, error)
	OwnByUser(collectionId string, userId string) bool
}

var (
	cache = gcache.New() // store code -> ttl \ $file_hash -> $file_url \ result_id-file_url -> $file url
)

type collectionService struct {
	log *glog.Logger
}

const CodeExists = 1

var (
	ErrCodeInvalid = errors.New("无效分享码")
	ErrVerifyCode  = errors.New("校验分享码失败")
)

type ServiceErr struct {
	inner error
	msg   string
}

func NewServiceErr(inner error, msg string) ServiceErr {
	return ServiceErr{inner, msg}
}

func (s ServiceErr) Error() string {
	switch s.inner.Error() {
	case "sql: no rows in result set":
		return ErrCodeInvalid.Error()
	default:
		return s.msg + ":" + s.inner.Error()
	}
}

var (
	_          collectionServiceI = &collectionService{}
	Collection                    = NewCollectionService()
)

func NewCollectionService() collectionServiceI {
	return &collectionService{
		log: g.Log("collection-service"),
	}
}

func (c *collectionService) FetchCollectionByCode(code string, userIp string) (*model.CollectionVO, error) {
	var (
		err error
	)
	codeInCache, err := cache.Get(nil, code)
	if err != nil {
		return nil, ErrVerifyCode
	}
	if codeInCache != nil && codeInCache.Int() != CodeExists {
		return nil, ErrCodeInvalid
	}
	var collection model.Collection
	err = dao.Collection.Ctx(nil).Where("code=? and validBefore > ?", code, gtime.Now()).Scan(&collection)
	if err != nil {
		return nil, NewServiceErr(err, "")
	}
	resultIdList, err := dao.UserResult.Ctx(nil).Fields("result_id").Where("collection_id=?", collection.Id).Array()
	if err != nil {
		return nil, NewServiceErr(err, "")
	}
	var results []*model.Result
	err = dao.Result.Ctx(nil).Where("id in (?)", resultIdList).Scan(&results)
	if err != nil {
		return nil, NewServiceErr(err, "")
	}

	// to collectionVO

	var ret model.CollectionVO
	convert.Transform(collection, &ret)
	var resList []*model.ResultVO
	for _, v := range results {
		var tmp model.ResultVO
		tmp.InvoiceDate = v.InvoiceDate.Format("Ymd")
		tmp.InvoiceCode = v.InvoiceCode
		tmp.CheckCode = v.CheckCode
		tmp.ParseType = v.ParseType
		tmp.SellerName = v.SellerName
		tmp.InvoiceType = v.InvoiceType
		tmp.InvoiceNumber = v.InvoiceNumber
		tmp.TotalAmount = v.TotalAmount
		// query and get file_url
		u, err := Result.getFileURL(int64(v.Id))
		if err != nil {
			c.log.Errorf(gctx.New(), "FetchCollectionByCode getFileURL result_id:%d failed:%s", v.Id, err)
		}
		tmp.FileURL = u
		resList = append(resList, &tmp)
	}
	ret.ResultList = resList
	return &ret, nil

}

func (r *collectionService) RegisterCollection(title, code string, userId string, userIp string, ttl int) (*model.Collection, error) {
	var args model.CreateCollectionServiceArgs
	ttlDuration := time.Second * time.Duration(ttl)
	args.Code = code
	args.UserId, _ = strconv.Atoi(userId)
	args.Ttl = ttl
	args.ValidBefore = gtime.Now().Add(ttlDuration)
	args.UserIp = userIp
	args.Title = title
	id, err := dao.Collection.Ctx(nil).InsertAndGetId(&args)
	if err == nil {
		// 校验放缓存
		err := cache.Set(nil, code, CodeExists, ttlDuration)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	var ret model.Collection
	err = dao.Collection.Ctx(nil).Where("id = ?", id).Scan(&ret)
	return &ret, err
}

// OwnByUser if a valid collection owned by user return true
func (r *collectionService) OwnByUser(collectionId string, userId string) bool {
	cnt1, err1 := dao.Collection.Ctx(nil).Count("user_id = ? and id = ? and validBefore > ?", userId, collectionId, gtime.Now())
	return cnt1 == 1 && err1 == nil
}
