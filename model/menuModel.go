package model

import "time"

type RestarauntMenu struct {
	Menu []Product `json:"menu"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Price       float32   `json:"price"`
	Type        string    `json:"type"`
	Ingredients string    `json:"ingredients"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   bool      `json:"is_deleted"`
	IDSupplier  int       `json:"id_supplier"`
}
