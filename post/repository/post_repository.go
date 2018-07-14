package repository

import (
	model "github.com/hack-boozer/boozer-api/post"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// postRepository post repository
type postRepository struct {
	Conn *gorm.DB
}

// NewPostRepository mount post repository
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		Conn: db,
	}
}

// PostRepository post repository interface
type PostRepository interface {
	GetByAccountID(accountID uuid.UUID) (*model.Post, error)
	Create(post *model.Post) (*model.Post, error)
}

func (m *postRepository) GetByAccountID(accountID uuid.UUID) (*model.Post, error) {
	post := model.Post{}
	err := m.Conn.Model(&post).Where("account_id = ?", accountID).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *postRepository) Create(post *model.Post) (*model.Post, error) {
	err := m.Conn.Create(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
