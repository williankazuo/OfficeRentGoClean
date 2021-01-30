package office

import "officerent/entity"

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
func (s *Service) GetAllOffices() ([]*entity.Office, error) {
	offices, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if len(offices) == 0 {
		return nil, entity.ErrNotFound
	}

	return offices, nil
}

// CreateOffice Create an office
func (s *Service) CreateOffice(title string, description string, people int, price float64) (entity.ID, error) {
	o, err := entity.NewOffice(title, description, people, price)
	if err != nil {
		return o.ID, err
	}

	return s.repo.Create(o)
}
