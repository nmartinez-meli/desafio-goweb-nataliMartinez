package tickets

import "context"

type Service interface {
	GetTotalTickets(c context.Context, destination string) (int, error)
	AverageDestination(c context.Context, destination string) (float64, error)
	// GetAll() ([]User, error)
	// CreateUser(nombre, apellido, email string, edad int64, altura float64, activo bool) (User, error)
	// GetUser(id int64) (User, error)
	// Update(nombre, apellido, email, fechaCreacion string, paramID, id, edad int64, altura float64, activo bool) (User, error)
	// UpdateField(nombre, apellido string, paramID int64) (User, error)
	// DeleteUser(paramID int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(c context.Context, destination string) (int, error) {
	totalDestination, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return -1, err
	}
	return len(totalDestination), nil
}

func (s *service) AverageDestination(c context.Context, destination string) (float64, error) {
	totalDestination, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return -1, err
	}
	total, err := s.repository.GetAll(c)
	if err != nil {
		return -1, err
	}
	avg := (float64(len(totalDestination)) / float64(len(total))) * 100
	return avg, nil
}
