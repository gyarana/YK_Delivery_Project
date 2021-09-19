package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Location     string `json:"location"`
	PhoneNumber  string `json:"phone_number"`
	Deleted      bool   `json:"deleted"`
}

type CurrentUser struct {
	ID int `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
