package account

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account account
type Account struct {
	ID        uuid.UUID  `form:"id" json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	Email     string     `form:"email" json:"email" xml:"email" gorm:"unique" name:"メールアドレス"`
	Password  string     `form:"password" json:"password" xml:"password" name:"パスワード"`
	Photo     string     `form:"photo" json:"photo" xml:"photo" name:"写真"`
	CreatedBy uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt *time.Time `json:"deletedAt" name:"削除日"`
}
