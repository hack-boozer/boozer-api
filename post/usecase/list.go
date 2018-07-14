package usecase

import "github.com/hack-boozer/boozer-api/post"

func (p *postUsecase) List() ([]*post.Media, error) {
	return p.postRepo.List()
}
