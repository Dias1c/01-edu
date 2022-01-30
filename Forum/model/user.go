package model

type User struct {
	Id       int    `json: "-"`
	Username string `json: "username"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func (u *User) Get() error {

}

func (u *User) Create() error {

}
