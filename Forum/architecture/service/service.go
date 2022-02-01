package service

import (
	"forum/architecture/repository"
	"forum/architecture/service/question"
	"forum/architecture/service/user"
)

type Service struct {
	User     IUser
	Question IQuestion
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     user.NewUserService(repo.User),
		Question: question.NewQuestionService(repo.Question),
	}
}
