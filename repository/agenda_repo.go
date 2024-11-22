package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type AgendaRepository interface {
	CreateAgendaRepo(ctx context.Context, agenda *domain.Agenda) *domain.Agenda
	FindAllAgendaRepo(ctx context.Context) []*domain.Agenda
	FindByIdAgendaRepo(ctx context.Context, id int64) (*domain.Agenda, error)
	CheckIDAgendaRepo(ctx context.Context, id int64) error
	UpdateAgendaRepo(ctx context.Context, agenda *domain.Agenda) *domain.Agenda

}