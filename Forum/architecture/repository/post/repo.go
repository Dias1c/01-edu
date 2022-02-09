package post

import (
	"fmt"

	"forum/architecture/helper"
	model "forum/architecture/models"
	"forum/configs"
)

func (p *PostRepo) Create(post *model.Post) error {
	query := `INSERT INTO posts (
		content, 
		user_id, 
		created_at
	) 
	VALUES (?, ?, ?);`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	encodedTime := helper.EncodeTime((*post).CreatedAt, configs.TimeFormatRFC3339)
	res, err := stmt.Exec((*post).Content, (*post).UserId, encodedTime)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	(*post).Id = lastId
	return nil
}

func (p *PostRepo) GetQuestion(postId int64) (*model.Post, error) {
	query := `SELECT id, content, user_id, created_at FROM posts WHERE id = ?`
	row := p.db.QueryRow(query, postId)

	post := model.Post{}
	decodedTime := ""

	err := row.Scan(&post.Id, &post.Content, &post.UserId, &decodedTime)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetQuestion: %w", err)
	}

	post.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetQuestion: %w", err)
	}

	return &post, nil
}

// takes a page number (1, 2, 3...)
func (p *PostRepo) GetQuestionAnswers(questionId int64, page int) ([]*model.Post, error) {
	query := `SELECT posts.*, sum(subquery.like_count) AS "like_count" FROM posts 
	INNER JOIN questions ON posts.id = questions.post_id WHERE questions.id = ? 
	`
	row := p.db.QueryRow(query, questionId)

	return nil, nil
}

func (p *PostRepo) GetMostLikedQuestion() (*model.Post, error) {
	return nil, nil
}
