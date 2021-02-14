package models

import (
	"officerent/entity"
	"time"
)

// OfficeCreate Model for Office Creation
type OfficeCreate struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	People      int     `json:"people"`
	Price       float64 `json:"price"`
	Country     int     `json:"country"`
	State       int     `json:"state"`
	City        int     `json:"city"`
	District    int     `json:"district"`
}

// OfficeRespose Model for Response of offices
type OfficeRespose struct {
	ID          entity.ID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	People      int        `json:"people"`
	Price       float64    `json:"price"`
	Country     string     `json:"country"`
	State       string     `json:"state"`
	City        string     `json:"city"`
	District    string     `json:"district"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
