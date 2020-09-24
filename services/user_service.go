package services

import (
	"fmt"
	"github.com/sswapnil2/bookstore_users_api/domain/users"
	"github.com/sswapnil2/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {

	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	result := users.User{
		Id: userId,
	}

	if err := result.Get(); err != nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", result.Id))
	}
	return &result, nil

}
