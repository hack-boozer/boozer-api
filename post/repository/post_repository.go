package repository

import (
	"database/sql"

	"github.com/hack-boozer/boozer-api/post"
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
	GetByAccountID(accountID uuid.UUID) (*post.Post, error)
	List() ([]*post.Media, error)
}

func (m *postRepository) GetByAccountID(accountID uuid.UUID) (*post.Post, error) {
	post := post.Post{}
	err := m.Conn.Model(&post).Where("account_id = ?", accountID).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *postRepository) List() ([]*post.Media, error) {
	rows, err := m.Conn.Table("posts").
		Select(`
			ac.nick_name,
			ac.photo,
			posts.comment,
			posts.rate,
			ph.file,
			posts.updated_at,
		`).
		Joins("LEFT JOIN accounts ac on ac.id = posts.account_id").
		Joins("LEFT JOIN photos ph on posts.id = ph.post_id").
		Where("posts.deleted_at is NULL").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := ListPosts(rows)
	return res, nil
}

// ListPosts list post
func ListPosts(rows *sql.Rows) []*post.Media {
	res := []*post.Media{}
	for rows.Next() {
		var p = post.Media{}
		rows.Scan(
			&p.User.NickName,
			&p.User.Photo,
			&p.Comment,
			&p.Rate,
			&p.Photo,
			&p.CreatedAt,
		)
		res = append(res, &p)
	}
	return res
}
