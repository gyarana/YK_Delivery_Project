package model

type Supplier struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AddressName string `json:"addressName"`
	IconURL     string `json:"iconURL"`
	TypeMarket  string `json:"typeMarket"`
}
