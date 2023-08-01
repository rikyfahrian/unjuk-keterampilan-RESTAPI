package usecase

import (
	"context"
	"errors"
	"restapi-altera/src/model"
	"restapi-altera/src/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoMock = &repository.RepositoryMock{Mock: mock.Mock{}}
var usecaseMock = usecase{repo: repoMock}
var ctx = context.Background()

func TestGetUserByIdFail(t *testing.T) {

	//testgagal

	repoMock.Mock.On("GetUserByID", ctx, "uuid").Return(nil, errors.New("not found"))

	user, err := usecaseMock.GetUserByID(ctx, "uuid")
	assert.Nil(t, user)
	assert.NotNil(t, err)

}

func TestGetUserByIDSucces(t *testing.T) {
	userSample := &model.User{
		Id:       "uuid2",
		Name:     "riky",
		Email:    "rikyfhrian@gmail.com",
		Password: "23",
		RackId:   "23",
	}

	repoMock.Mock.On("GetUserByID", ctx, "uuid2").Return(userSample, nil)

	user, _ := usecaseMock.GetUserByID(ctx, "uuid2")

	assert.Equal(t, userSample.Name, user.Name)
}
