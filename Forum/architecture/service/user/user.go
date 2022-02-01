package user

import "forum/architecture/repository"

type UserService struct {
	repo repository.IUser
}

func NewUserService(repo repository.IUser) *UserService {
	return &UserService{repo}
}
