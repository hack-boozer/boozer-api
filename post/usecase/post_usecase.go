package usecase

import (
	"database/sql"

	"github.com/hack-boozer/boozer-api/photo"
	photoRepo "github.com/hack-boozer/boozer-api/photo/repository"
	"github.com/hack-boozer/boozer-api/post"
	postRepo "github.com/hack-boozer/boozer-api/post/repository"
	uuid "github.com/satori/go.uuid"
)

type postUsecase struct {
	postRepo  postRepo.PostRepository
	photoRepo photoRepo.PhotoRepository
}

// NewPostUsecase mount post usecase
func NewPostUsecase(postRepo postRepo.PostRepository, photoRepo photoRepo.PhotoRepository) PostUsecase {
	return &postUsecase{
		postRepo:  postRepo,
		photoRepo: photoRepo,
	}
}

// PostUsecase post usecase interface
type PostUsecase interface {
	Create(param *post.Create) (*post.Post, error)
}

func formatPostPayload(param *post.Create) (*post.Post, []*photo.Photo) {
	postReq := &post.Post{
		ID:        uuid.NewV4(),
		AccountID: param.AccountID,
		Comment:   param.Comment,
		Rate:      sql.NullFloat64{Float64: param.Rate, Valid: true},
		UpdatedBy: param.AccountID,
	}
	photoReq := []*photo.Photo{}
	for _, p := range param.Photos {
		var req = &photo.Photo{
			PostID: postReq.ID,
			File:   p,
		}
		photoReq = append(photoReq, req)
	}
	return postReq, photoReq
}
