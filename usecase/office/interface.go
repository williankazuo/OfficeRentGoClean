package office

import (
	"officerent/api/models"
	"officerent/entity"
)

// Repository Office repository interface
type Repository interface {
	Create(o *entity.Office) (entity.ID, error)
	List() ([]*entity.Office, error)
	Get(id string) (*models.OfficeRespose, error)
}

// UseCase Office usecase interface
type UseCase interface {
	GetAllOffices() ([]*entity.Office, error)
	GetOfficeDetail(id string) (*models.OfficeRespose, error)
	CreateOffice(title string, description string, people int, price float64, country, state, city, district int) (entity.ID, error)
}
