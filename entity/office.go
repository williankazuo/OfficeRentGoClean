package entity

import "time"

// Office entity
type Office struct {
	ID          ID
	Title       string
	Description string
	People      int
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewOffice Creates a new Office entity
func NewOffice(title string, description string, people int, price float64) (*Office, error) {
	o := &Office{
		ID:          NewID(),
		Title:       title,
		Description: description,
		People:      people,
		Price:       price,
		CreatedAt:   time.Now(),
	}

	err := o.validate()
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Office) validate() error {
	if o.Title == "" || o.People <= 0 || o.Price <= 0 {
		return ErrInvalidEntity
	}
	return nil
}
