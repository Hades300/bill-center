// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserResultDao is the data access object for table user_result.
type UserResultDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns UserResultColumns // Columns contains all the column names of Table for convenient usage.
}

// UserResultColumns defines and stores column names for table user_result.
type UserResultColumns struct {
	UserId    string // 用户id                
    FileHash  string // 文件哈希              
    FileUrl   string // 若解析失败，上传文件  
    ResultId  string // 结果id
}

//  userResultColumns holds the columns for table user_result.
var userResultColumns = UserResultColumns{
	UserId:   "user_id",     
            FileHash: "file_hash",   
            FileUrl:  "file_url",    
            ResultId: "result__id",
}

// NewUserResultDao creates and returns a new DAO object for table data access.
func NewUserResultDao() *UserResultDao {
	return &UserResultDao{
		Group:   "default",
		Table:   "user_result",
		Columns: userResultColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserResultDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserResultDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}