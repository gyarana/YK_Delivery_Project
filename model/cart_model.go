package model

import "time"

type CartProducts struct {
	CartID    int `json:"cartID"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	ID          int            `json:"id"`
	Products    []CartProducts `json:"products"`
	CreatedDate time.Time      `json:"created_date"`
	UpdatedDate time.Time      `json:"updated_date"`
	DeletedDate time.Time      `json:"deleted_date"`
	IsDeleted   bool           `json:"is_deleted"`
}

type CartRequest struct {
	ID int `json:"id"`
}
