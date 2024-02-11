package response

type UserResponse struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
}
