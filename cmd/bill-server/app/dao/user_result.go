// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hades300/bill-center/cmd/bill-server/app/dao/internal"
)

// userResultDao is the data access object for table user_result.
// You can define custom methods on it to extend its functionality as you wish.
type userResultDao struct {
	*internal.UserResultDao
}

var (
	// UserResult is globally public accessible object for table user_result operations.
	UserResult = userResultDao{
		internal.NewUserResultDao(),
	}
)

// Fill with you ideas below.
