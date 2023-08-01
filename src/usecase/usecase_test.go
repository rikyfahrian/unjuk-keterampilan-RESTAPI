package usecase

import (
	"context"
	"errors"
	"restapi-altera/src/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoMock = &repository.RepositoryMock{Mock: mock.Mock{}}
var usecaseMock = usecase{repo: repoMock}

func TestGetUserById(t *testing.T) {

	ctx := context.Background()

	repoMock.Mock.On("GetUserById", ctx, "uuid").Return(nil, errors.New("not found"))

	user, err := usecaseMock.GetUserByID(ctx, "uuid")
	assert.Nil(t, user)
	assert.NotNil(t, err)

}
