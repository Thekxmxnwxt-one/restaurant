package models

type Menu struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageURL    *string `json:"image_url"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
	Category    *string `json:"category"`
	Available   bool    `json:"available"`
}
