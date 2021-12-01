// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectionDao is the data access object for table collection.
type CollectionDao struct {
	Table   string            // Table is the underlying table name of the DAO.
	Group   string            // Group is the database configuration group name of current DAO.
	Columns CollectionColumns // Columns contains all the column names of Table for convenient usage.
}

// CollectionColumns defines and stores column names for table collection.
type CollectionColumns struct {
	Title       string // 集合标题
	UserId      string // 外键 留用
	Id          string // 主键
	UserIp      string // 用户IP
	GmtCreated  string // 创建UTC时间
	Code        string // 集合密码
	GmtModified string // 更新UTC时间
	Ttl         string // 有效期（单位秒）
	ValidBefore string // 失效时间(后端主动控制）
}

//  collectionColumns holds the columns for table collection.
var collectionColumns = CollectionColumns{
	Title:       "title",
	UserId:      "user_id",
	Id:          "id",
	UserIp:      "user_ip",
	GmtCreated:  "gmtCreated",
	Code:        "code",
	GmtModified: "gmtModified",
	Ttl:         "ttl",
	ValidBefore: "validBefore",
}

// NewCollectionDao creates and returns a new DAO object for table data access.
func NewCollectionDao() *CollectionDao {
	return &CollectionDao{
		Group:   "default",
		Table:   "collection",
		Columns: collectionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectionDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
