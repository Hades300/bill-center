// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/hades300/bill-center/cmd/bill-server/app/dao/internal"
)

// collectionDao is the data access object for table collection.
// You can define custom methods on it to extend its functionality as you wish.
type collectionDao struct {
	*internal.CollectionDao
}

var (
	// Collection is globally public accessible object for table collection operations.
	Collection = collectionDao{
		internal.NewCollectionDao(),
	}
)

// Fill with you ideas below.
