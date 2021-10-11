package model

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Location     string    `json:"location"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	DeletedDate  time.Time `json:"deleted_date"`
	IsDeleted    bool      `json:"is_deleted"`
}

type CurrentUser struct {
	ID int `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
