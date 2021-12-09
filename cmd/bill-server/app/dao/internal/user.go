// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	Table   string      // Table is the underlying table name of the DAO.
	Group   string      // Group is the database configuration group name of current DAO.
	Columns UserColumns // Columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id          string // 主键
	Phone       string // 手机号
	Email       string // 邮箱
	GmtCreated  string //
	GmtModified string //
	Nickname    string // 昵称
	Password    string //
}

//  userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:          "id",
	Phone:       "phone",
	Email:       "email",
	GmtCreated:  "gmtCreated",
	GmtModified: "gmtModified",
	Nickname:    "nickname",
	Password:    "password",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		Group:   "default",
		Table:   "user",
		Columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
