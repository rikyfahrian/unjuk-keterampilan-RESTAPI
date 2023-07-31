package usecase

import (
	"context"
	"errors"
	"os"
	"restapi-altera/src/model"

	"github.com/golang-jwt/jwt/v4"
)

func (u *usecase) GetUserByID(ctx context.Context, cookie string) (*model.User, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	user, err := u.repo.GetUserByID(ctx, claims.Issuer)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}
