package usecase

import (
	"context"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/repository"
)

type Usecase interface {
	Register(ctx context.Context, user request.UserRegister) error
	Login(ctx context.Context, user request.UserLogin) (string, error)
	GetUserByID(ctx context.Context, cookie string) (*model.User, error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}
