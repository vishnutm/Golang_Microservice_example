package services

import (
	"pack/domain/users"
	"pack/utils/errors"
)

func CreateUser(user users.User) (*users.User,*errors.RestErr)  {

	return &user, nil
	
}