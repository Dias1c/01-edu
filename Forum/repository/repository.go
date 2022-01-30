package repository

type User interface {
}

type Question interface {
}

type Repository struct {
	User     *User
	Question *Question
}

func NewRepository() *Repository {
	return &Repository{}
}
