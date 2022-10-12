package models

type UserRequest struct {
	Name     string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type UserResp struct {
	Name  string `json:"username"`
	Token string `json:"token"`
}
