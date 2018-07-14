package migrations

import (
	"github.com/hack-boozer/boozer-api/account"
	"github.com/hack-boozer/boozer-api/db"
	"github.com/jinzhu/gorm"
)

// Execute マイグレーション実行
func Execute() {
	odb := db.ConnectDB()
	Migrate(odb)
	odb.Close()
}

// Migrate マイグレーション
func Migrate(odb *gorm.DB) {
	odb.AutoMigrate(
		&account.Account{},
	)
}

// DropTable テーブル削除
func DropTable(odb *gorm.DB) {
	odb.DropTableIfExists(
		&account.Account{},
	)
}
