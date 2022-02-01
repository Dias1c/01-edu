package service

import (
	model "forum/architecture/models"
)

type IUser interface {
	Create(user *model.User) error
	Update(user *model.User) error
	DeleteByID(id int) error

	GetByID(id int) (*model.User, error)
	// GetByNickname(nickname string) (*model.User, error)
	// GetAll(from, offset int) error

	// CanLogin(user *model.User) error
}
