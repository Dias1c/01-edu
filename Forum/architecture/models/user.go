package models

import "time"

// User -
type User struct {
	Id          int
	Nickname    string
	Fistname    string
	Lastname    string
	Password    string
	CreatedTime time.Time
	// PHOTO?
}
