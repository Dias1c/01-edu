package repository

import (
	model "forum/architecture/models"
)

type IQuestion interface {
	Create(user *model.Question) error
	Update(user *model.Question) error
	GetByID(id int) (*model.Question, error)
	DeleteByID(id int) error
}
