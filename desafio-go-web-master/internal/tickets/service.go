package tickets

import (
	"context"

	"github.com/juanimeli/desafio-goweb-juanguglielmone/desafio-go-web-master/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
	GetTicketsByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type service struct {
	respoitory Repository
}

func NewService(r Repository) Service {
	return &service{
		respoitory: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	t, err := s.respoitory.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {

	t, err := s.respoitory.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(t), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {

	totalT, err := s.respoitory.GetAll(ctx)
	if err != nil {
		return 0.0, err
	}
	tDestination, err := s.respoitory.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0.0, err
	}

	prom := float64(len(tDestination)) / float64(len(totalT)) * 100

	return prom, nil

}

func (s *service) GetTicketsByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	t, err := s.respoitory.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return t, nil
}
