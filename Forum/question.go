package forum

type Question struct {
	Id     int    `json: "id"`
	Title  string `json: "title"`
	Text   string `json: "text"`
	UserId int    `json: "title"`
}
