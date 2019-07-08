package models

import "fmt"

type User struct {
	Id   string
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User's id : %s, name : %s", u.Id, u.Name)
}
