package repository

import (
	"database/sql"

	"forum/architecture/repository/question"
	"forum/architecture/repository/user"
)

//Repository ..
type Repository struct {
	User     IUser
	Question IQuestion
}

//NewRepo ..
func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:     user.NewUserRepo(db),
		Question: question.NewQuestionRepo(db),
	}
}
