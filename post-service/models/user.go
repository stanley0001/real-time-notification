package models

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
