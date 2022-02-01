package models

type Question struct {
	Id    int
	Title string
	Text  string
	// Tags     []*Tag
	// Answers  []*Answer
	UserId int
}
