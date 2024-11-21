package impl_repo

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"gorm.io/gorm"
)

func NewAgendaRepostiory(DB *gorm.DB) repository.AgendaRepository {
	return &AgendaRepositoryImpl{
		DB: DB,
	}
}

type AgendaRepositoryImpl struct {
	*gorm.DB
}

func (repo *AgendaRepositoryImpl)CreateAgendaRepo(ctx context.Context, agenda *domain.Agenda) *domain.Agenda{
	err := repo.DB.WithContext(ctx).Create(&agenda).Error
	exception.PanicLogging(err)
	return agenda
}

func (repo *AgendaRepositoryImpl)FindAllAgendaRepo(ctx context.Context) []*domain.Agenda {
	var agendas []*domain.Agenda
	err := repo.DB.WithContext(ctx).Find(&agendas).Error
	exception.PanicLogging(err)
	return agendas
}