package question

import (
	"forum/architecture/repository"
)

type QuestionService struct {
	repo repository.IQuestion
}

func NewQuestionService(repo repository.IQuestion) *QuestionService {
	return &QuestionService{repo}
}
