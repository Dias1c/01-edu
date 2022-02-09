package models

import "time"

type Comment struct {
	Id        int64
	Content   string
	UserId    int64
	CreatedAt time.Time
	// PostId  int
}

type ICommentRepo interface {
	Create(comment *Comment) error
	// Delete(id int) error
	GetPostComments(postId int64) ([]*Comment, error)
}
