package model

type User struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Location    string `json:"location"`
	PhoneNumber string `json:"phone_number"`
	Deleted     bool   `json:"deleted"`
}
