package model

import "time"

type Order struct {
	OrderID     int32     `json:"id"`
	UserID      int32     `json:"user_id"`
	CartID      int32     `json:"cart_id"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   bool      `json:"is_deleted"`
}

type OrderRequestID struct {
	ID int32 `json:"id"`
}

type OrderRequest struct {
	OrderID int32  `json:"id"`
	UserID  int32  `json:"user_id"`
	CartID  int32  `json:"cart_id"`
	Status  string `json:"status"`
}
