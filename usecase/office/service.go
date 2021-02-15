package office

import (
	"officerent/api/models"
	"officerent/entity"
)

// Service Office usecase
type Service struct {
	repo Repository
}

// NewService Create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// GetAllOffices Get all the offices
func (s *Service) GetAllOffices() ([]*models.OfficeRespose, error) {
	offices, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if len(offices) == 0 {
		return nil, entity.ErrNotFound
	}

	return offices, nil
}

// GetOfficeDetail Get an office detail
func (s *Service) GetOfficeDetail(id string) (*models.OfficeRespose, error) {
	office, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if office == nil {
		return nil, entity.ErrNotFound
	}

	return office, nil
}

// CreateOffice Create an office
func (s *Service) CreateOffice(title string, description string, people int, price float64, country, state, city, district int) (entity.ID, error) {
	o, err := entity.NewOffice(title, description, people, price, country, state, city, district)
	if err != nil {
		return o.ID, err
	}

	return s.repo.Create(o)
}
