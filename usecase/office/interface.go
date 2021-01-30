package office

import "officerent/entity"

// Repository Office repository interface
type Repository interface {
	Create(o *entity.Office) (entity.ID, error)
	List() ([]*entity.Office, error)
}

// UseCase Office usecase interface
type UseCase interface {
	GetAllOffices() ([]*entity.Office, error)
	CreateOffice(title string, description string, people int, price float64) (entity.ID, error)
}
