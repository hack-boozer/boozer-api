package repository

import (
	"github.com/hack-boozer/boozer-api/account"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// accountRepository account repository
type accountRepository struct {
	Conn *gorm.DB
}

// NewAccountRepository mount account repository
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

// AccountRepository account repository interface
type AccountRepository interface {
	GetByID(id uuid.UUID) (*account.Account, error)
}

func (m *accountRepository) GetByID(id uuid.UUID) (*account.Account, error) {
	account := account.Account{}
	err := m.Conn.Model(&account).Where("id = ?", id).Find(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
