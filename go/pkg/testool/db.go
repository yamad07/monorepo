package testool

import (
	"testing"

	"github.com/yamad07/monorepo/go/pkg/mysql"
	"gorm.io/gorm"
)

func PrepareDB(t *testing.T, name string, seeds []interface{}) (*gorm.DB, func()) {
	db, err := mysql.NewTest(name)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			t.Fatal(err)
		}
	}
	return db, func() { sqldb, _ := db.DB(); sqldb.Close() }
}

func CountRecord(t *testing.T, db *gorm.DB, table string) int {
	var cnt int64
	if err := db.Table(table).Count(&cnt).Error; err != nil {
		t.Fatal(err)
	}
	return int(cnt)
}
