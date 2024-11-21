package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type AgendaService interface {
	CreateAgendaService(ctx context.Context, agenda *web.AgendaCreateOrUpdate) *web.AgendaCreateOrUpdate
	FindAllAgendaService(ctx context.Context) []*web.AgendaCreateOrUpdate
	FindByIdAgendaService(ctx context.Context, id int64) *web.AgendaModel
}