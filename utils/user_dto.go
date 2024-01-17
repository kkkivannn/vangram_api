package utils

type UserDto struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}
