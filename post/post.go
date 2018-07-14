package post

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Post post
type Post struct {
	ID        uuid.UUID       `json:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	AccountID uuid.UUID       `json:"accountId" sql:"type:uuid" name:"アカウントID"`
	Comment   string          `json:"comment" name:"コメント"`
	Rate      sql.NullFloat64 `json:"rate" name:"レート"`
	CreatedBy uuid.UUID       `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID       `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time       `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time       `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time      `json:"deletedAt" name:"削除日"`
}
