package model

import "time"

type Suppliers struct {
	Restaurants []Restaurant `json:"suppliers"`
}

type Restaurant struct {
	Id           int          `json:"id"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	WorkingHours WorkingHours `json:"workingHours"`
	CreatedDate  time.Time    `json:"created_date"`
	UpdatedDate  time.Time    `json:"updated_date"`
	DeletedDate  time.Time    `json:"deleted_date"`
	IsDeleted    bool         `json:"is_deleted"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type SupplierRequestID struct {
	ID int `json:"id"`
}

type RestaurantParse struct {
	Id           int          `json:"id"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	WorkingHours WorkingHours `json:"workingHours"`
}
