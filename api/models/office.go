package models

// OfficeCreate Model for Office Creation
type OfficeCreate struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	People      int     `json:"people"`
	Price       float64 `json:"price"`
}
