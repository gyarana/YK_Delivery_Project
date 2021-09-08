package model

type Order struct {
	Id       int32     `json:"id"`
	Entities []Product `json:"entities"`
	Status   string    `json:"status"`
	Adress   string    `json:"adress"`
}
