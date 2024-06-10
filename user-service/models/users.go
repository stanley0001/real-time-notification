package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}
