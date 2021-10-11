package model

import "time"

type Supliers struct {
	Restaurants []Restaurant `json:"suppliers"`
}

type Restaurant struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   bool      `json:"is_deleted"`
}
