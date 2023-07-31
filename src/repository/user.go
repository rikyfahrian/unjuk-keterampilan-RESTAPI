package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/model/response"
	"time"
)

func (r *repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {

	var user *model.User

	err := r.config.Database().WithContext(ctx).Where("id = ? ", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUserAllShoes(ctx context.Context, id string) (*response.UserShoes, error) {

	var resp *response.UserShoes

	//caching
	r.key = fmt.Sprintf("allshoesUser%s", id)
	data, err := r.config.Redis().Get(ctx, r.key)
	if err != nil {
		var user *model.User
		var AllShoes []*model.Shoes

		tx := r.config.Database().WithContext(ctx).Begin()
		err := tx.Where("id = ?", id).First(&user).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.Where("shoes_id = ?", user.RackId).Find(&AllShoes).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		tx.Commit()

		resp = &response.UserShoes{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			RackShoes: AllShoes,
		}

		err = r.config.Redis().Set(ctx, r.key, resp)
		if err != nil {
			return nil, err
		}

		return resp, nil

	}

	err = json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (r *repository) GetUserShoesByNoId(ctx context.Context, noId int, id string) (*response.UserShoes, error) {

	var user *model.User

	tx := r.config.Database().WithContext(ctx).Begin()

	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var Shoes *model.Shoes

	err = tx.Where("shoes_id = ? AND no_id = ?", user.RackId, noId).First(&Shoes).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	resp := &response.UserShoes{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		RackShoes: Shoes,
	}

	return resp, nil

}

func (r *repository) AddShoes(ctx context.Context, shoes *request.ReqUserShoes, id string) error {

	var user *model.User

	tx := r.config.Database().WithContext(ctx).Begin()

	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	payload := &model.Shoes{
		ShoesId: user.RackId,
		Name:    shoes.Name,
		Size:    shoes.Size,
		AddAt:   time.Now().Format(time.RFC822),
	}

	err = tx.Create(payload).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	r.config.Redis().Delete(ctx, r.key)
	return nil

}

func (r *repository) UpdateShoesByNoId(ctx context.Context, noId int, id string, payload *request.UpdateUserShoes) error {

	var user *model.User

	tx := r.config.Database().WithContext(ctx).Begin()

	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	payloadNew := &model.Shoes{
		Name:      payload.Name,
		Size:      payload.Size,
		UpdatedAt: time.Now().Format(time.RFC822),
	}

	err = tx.Where("no_id = ? AND shoes_id = ?", noId, user.RackId).Updates(payloadNew).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}

func (r *repository) DeleteShoesByNoId(ctx context.Context, id string, noId int) error {

	var user *model.User

	tx := r.config.Database().WithContext(ctx).Begin()

	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	result := tx.Where("no_id = ? AND shoes_id = ?", noId, user.RackId).Delete(&model.Shoes{})
	if result.Error != nil {
		tx.Rollback()
		return errors.New("failed to deleted data")
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("failed to deleted data")
	}

	tx.Commit()

	return nil
}
