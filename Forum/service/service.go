package service

import "forum/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
