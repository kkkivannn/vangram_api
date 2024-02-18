package user

type User struct {
	ID      *int    `json:"id,omitempty"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
}

type RequestUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
