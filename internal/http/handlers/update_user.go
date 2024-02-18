package handlers

type RequestUpdateUser struct {
	ID      int     `json:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
}
