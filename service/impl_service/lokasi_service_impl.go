package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

func NewLokasiServiceImpl(lokasiRepository *repository.LokasiRepository) service.LokasiService {
	return &LokasiServiceImpl{
		LokasiRepository: *lokasiRepository,
	}
}

type LokasiServiceImpl struct {
	repository.LokasiRepository
}

func (service *LokasiServiceImpl) CreateLokasiService(ctx context.Context, lokasi *web.LokasiCreateOrUpdate) *web.LokasiCreateOrUpdate {
	configuration.Validate(lokasi)
	lokasis := &domain.Lokasi{
		Nama: lokasi.Nama,
		Alamat: lokasi.Alamat,
	}
	service.LokasiRepository.CreateLokasiRepo(ctx, lokasis)
	return lokasi
}