package models

// OfficeCreate Model for Office Creation
type OfficeCreate struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	People      int     `json:"people"`
	Price       float64 `json:"price"`
	Country     string  `json:"country"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	District    string  `json:"district"`
}
