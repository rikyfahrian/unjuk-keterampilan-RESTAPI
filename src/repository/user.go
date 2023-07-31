package repository

import (
	"context"
	"restapi-altera/src/model"
)

func (r *repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {

	var user *model.User

	err := r.config.Database().WithContext(ctx).Where("id = ? ", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
