package service

import (
	"strconv"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hades300/bill-center/cmd/bill-server/app/dao"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
)

type collectionServiceI interface {
	FetchCollectionByCode(code string, userIp string) (*model.CollectionVO, error)
	RegisterCollection(code string, userId string, userIp string, ttl int) error
}

type collectionService struct{}

var (
	_          collectionServiceI = &collectionService{}
	Collection                    = NewCollectionService()
)

func NewCollectionService() collectionServiceI {
	return &collectionService{}
}

func (c *collectionService) FetchCollectionByCode(code string, userIp string) (*model.CollectionVO, error) {
	var collection model.Collection
	err := dao.Collection.Ctx(nil).Where("code=?", code).Scan(&collection)
	if err != nil {
		return nil, err
	}
	resultIdList, err := dao.UserResult.Ctx(nil).Fields("result_id").Where("collection_id=?", collection.Id).Array()
	if err != nil {
		return nil, err
	}
	var results []*model.Result
	err = dao.Result.Ctx(nil).Where("id in (?)", resultIdList).Scan(&results)
	if err != nil {
		return nil, err
	}
	collectionVO := model.ToCollectionVO(&collection, results)
	return collectionVO, nil
}

func (r *collectionService) RegisterCollection(code string, userId string, userIp string, ttl int) error {
	var collection model.Collection
	collection.Code = code
	collection.UserId, _ = strconv.Atoi(userId)
	collection.Ttl = ttl
	collection.ValidBefore = gtime.Now().Add(time.Second * time.Duration(ttl))
	collection.UserIp = userIp
	_, err := dao.Collection.Ctx(nil).Insert(&collection)
	return err
}
