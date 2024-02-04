package utils

type Request struct {
	Id      *int    `json:"id" db:"id"`
	Name    *string `json:"name" db:"name"`
	Surname *string `json:"surname" db:"surname"`
}

type InputUser struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}
