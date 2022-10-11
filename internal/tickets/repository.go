package tickets

import (
	"context"
	"fmt"
	"https://github.com/TomasCuelloRodriguezMl/desafio-goweb-tomascuello/internal/domains"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domains.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domains.Ticket, error)
}

type repository struct {
	db []domains.Ticket
}

func NewRepository(db []domains.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domains.Ticket, error) {

	if len(r.db) == 0 {
		return []domains.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domains.Ticket, error) {

	var ticketsDest []domains.Ticket

	if len(r.db) == 0 {
		return []domains.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}
