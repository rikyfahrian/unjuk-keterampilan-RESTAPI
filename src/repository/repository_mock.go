package repository

import (
	"context"
	"errors"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/model/response"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (rm *RepositoryMock) Register(ctx context.Context, user *model.User) error {
	panic("implement me")
}
func (rm *RepositoryMock) Login(ctx context.Context, user *request.UserLogin) (*model.User, error) {
	panic("implement me")
}

func (rm *RepositoryMock) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	arguments := rm.Mock.Called(ctx, id)

	if arguments.Get(0) == nil {
		return nil, errors.New("not found")
	}

	user := arguments.Get(0).(model.User)

	return &user, nil

}

func (rm *RepositoryMock) GetUserAllShoes(ctx context.Context, id string) (*response.UserShoes, error) {
	panic("implement me")
}
func (rm *RepositoryMock) GetUserShoesByNoId(ctx context.Context, noId int, id string) (*response.UserShoes, error) {
	panic("implement me")
}
func (rm *RepositoryMock) AddShoes(ctx context.Context, shoes *request.ReqUserShoes, id string) error {
	panic("implement me")
}
func (rm *RepositoryMock) UpdateShoesByNoId(ctx context.Context, noId int, id string, payload *request.UpdateUserShoes) error {
	panic("implement me")
}
func (rm *RepositoryMock) DeleteShoesByNoId(ctx context.Context, id string, noId int) error {
	panic("implement me")
}
