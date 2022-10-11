package tickets

import (
	"context"
	"desafio-goweb-tomascuello/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)

	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}

// AverageDestination implements Service
func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	panic("unimplemented")
}

// GetAll implements Service
func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// GetTicketByDestination implements Service
func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
