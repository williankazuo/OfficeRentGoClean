package office

import "officerent/entity"

type Service struct {
	repo Repository
}

func (s *Service) GetAllOffices() ([]entity.Office, error) {
	offices, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if len(offices) == 0 {
		return nil, entity.ErrNotFound
	}

	return offices, nil
}

func (s *Service) CreateOffice(title string, description string, people int, price float64) (entity.ID, error) {
	o, err := entity.NewOffice(title, description, people, price)
	if err != nil {
		return o.ID, err
	}

	return s.repo.Create(o)
}
