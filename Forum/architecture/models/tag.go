package models

type Tag struct {
	Id    int64
	Name  string
	Count int
}

type ITagRepo interface {
	Create(tag *Tag) error
	GetTags(page int) ([]*Tag, error)
}

type ITagService interface {
	GetTags(page int) ([]*Tag, error)
}
