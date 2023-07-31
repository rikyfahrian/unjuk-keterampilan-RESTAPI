package repository

import (
	"context"
	"restapi-altera/config"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
)

type Repository interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *request.UserLogin) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type repository struct {
	config config.Config
}

func NewRepository(config config.Config) Repository {
	return &repository{
		config: config,
	}
}
