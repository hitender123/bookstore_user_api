package services

import (
	"vijayapi/users_api/model/users"
	"vijayapi/users_api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(uid int64) (*users.User, *errors.RestError)
	CreateUser(user users.User) (*users.User, *errors.RestError)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError)
	DeleteUser(userId int64) *errors.RestError
}

func (s *usersService) GetUser(uid int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: uid}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.Address != "" {
			current.Address = user.Address
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.Address = user.Address
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	dao := &users.User{Id: userId}
	return dao.Delete()
}
