package models

import "time"

type Post struct {
	Id        int64
	Content   string
	UserId    int64
	CreatedAt time.Time
	Comments  []*Comment
	Likes     int
	Dislikes  int
	Rate      int
}

type IPostRepo interface {
	Create(post *Post) error
	// Delete(id int) error
	GetQuestion(postId int64) (*Post, error)
	GetQuestionAnswers(questionId int64) ([]*Post, error)
	GetMostLikedQuestion() (*Post, error)
	GetRate(post *Post) error
}

type IPostService interface {
	Create(post *Post) error
	Like(like *Like) error
	Dislike(like *Like) error
	Unlike(like *Like) error
	Comment(comment *Comment) error
}
