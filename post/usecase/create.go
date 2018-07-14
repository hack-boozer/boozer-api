package usecase

import "github.com/hack-boozer/boozer-api/post"

func (p *postUsecase) Create(param *post.Create) (*post.Post, error) {
	postReq, photoReq := formatPostPayload(param)
	postReq.CreatedBy = postReq.UpdatedBy

	res, err := p.postRepo.Create(postReq)
	if err != nil {
		return nil, err
	}
	for _, img := range photoReq {
		_, err = p.photoRepo.Create(img)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
