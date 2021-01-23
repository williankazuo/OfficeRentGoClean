package office

import "officerent/entity"

type Repository interface {
	Create(o *entity.Office) (entity.ID, error)
	List() ([]entity.Office, error)
}

type UseCase interface {
	GetAllOffices() (*[]entity.Office, error)
	CreateOffice(title string, description string, people int, price float64) (entity.ID, error)
}
