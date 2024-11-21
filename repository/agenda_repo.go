package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type AgendaRepository interface {
	CreateAgendaRepo(ctx context.Context, agenda *domain.Agenda) *domain.Agenda
}