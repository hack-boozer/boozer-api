package usecase

import (
	"github.com/hack-boozer/boozer-api/account"
	"github.com/hack-boozer/boozer-api/account/repository"
	uuid "github.com/satori/go.uuid"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

// AccountUsecase account usecase interface
type AccountUsecase interface {
	GetAccount(id uuid.UUID) (*account.Media, error)
}

// NewAccountUsecase mount account usecase
func NewAccountUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
	}
}

func (a *accountUsecase) GetAccount(id uuid.UUID) (*account.Media, error) {
	user, err := a.accountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	res := &account.Media{
		ID:       user.ID,
		NickName: user.NickName,
		Email:    user.Email,
		Photo:    user.Photo,
	}
	return res, nil
}
