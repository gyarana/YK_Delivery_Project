package model

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}
