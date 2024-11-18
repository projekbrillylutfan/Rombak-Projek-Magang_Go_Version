package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

func NewBupatiServiceImpl(bupatiRepository *repository.BupatiRepository) service.BupatiService {
	return &BupatiServiceImpl{
		BupatiRepository: *bupatiRepository,
	}
}

type BupatiServiceImpl struct {
	repository.BupatiRepository
}

func (service *BupatiServiceImpl) CreateBupatiService(ctx context.Context, bupati *web.BupatiCreateOrUpdate) *web.BupatiCreateOrUpdate {
	configuration.Validate(bupati)
	bupatis := &domain.Bupati{
		Nama: bupati.Nama,
		PeriodeJabatan: bupati.PeriodeJabatan,
	}

	service.BupatiRepository.CreateBupatiRepo(ctx, bupatis)
	return bupati
}