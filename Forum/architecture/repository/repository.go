package repository

import (
	"database/sql"
	"forum/architecture/models"
	"forum/architecture/repository/comment"
	"forum/architecture/repository/like"
	"forum/architecture/repository/post"
	"forum/architecture/repository/question"
	"forum/architecture/repository/session"
	"forum/architecture/repository/tag"
	"forum/architecture/repository/user"
)

type Repository struct {
	User     models.IUserRepo
	Session  models.ISessionRepo
	Question models.IQuestionRepo
	Post     models.IPostRepo
	Comment  models.ICommentRepo
	Like     models.ILikeRepo
	Tag      models.ITagRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:     user.NewUserRepo(db),
		Session:  session.NewSessionRepo(db),
		Question: question.NewQuestionRepo(db),
		Post:     post.NewPostRepo(db),
		Comment:  comment.NewCommentRepo(db),
		Like:     like.NewLikeRepo(db),
		Tag:      tag.NewTagRepo(db),
	}
}
