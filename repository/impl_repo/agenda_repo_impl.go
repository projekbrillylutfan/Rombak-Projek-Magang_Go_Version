package impl_repo

import (
	"context"
	"errors"

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

func (repo *AgendaRepositoryImpl)FindByIdAgendaRepo(ctx context.Context, id int64) (*domain.Agenda, error) {
	var agenda *domain.Agenda
	result := repo.DB.WithContext(ctx).Unscoped().Preload("Bupati").Preload("Lokasi").Preload("JenisAcara").Where("id_agenda", id).First(&agenda)
	if result.RowsAffected == 0 {
		return &domain.Agenda{}, errors.New("agenda not found")
	}
	return agenda, nil
}

func (repo *AgendaRepositoryImpl)CheckIDAgendaRepo(ctx context.Context, id int64) (*domain.Agenda, error) {
	var agenda *domain.Agenda
	result := repo.DB.WithContext(ctx).Unscoped().Where("id_agenda = ?", id).First(&agenda)
	if result.RowsAffected == 0 {
		return &domain.Agenda{}, errors.New("agenda not found")
	}
	return agenda, nil
}

func (repo *AgendaRepositoryImpl)UpdateAgendaRepo(ctx context.Context, agenda *domain.Agenda) *domain.Agenda {
	err := repo.DB.WithContext(ctx).Where("id_agenda = ?", agenda.IDAgenda).Updates(&agenda).Error
	exception.PanicLogging(err)
	return agenda
}

func (repo *AgendaRepositoryImpl)DeleteAgendaRepo(ctx context.Context, agenda *domain.Agenda){
	err := repo.DB.WithContext(ctx).Delete(&agenda).Error
	exception.PanicLogging(err)
}