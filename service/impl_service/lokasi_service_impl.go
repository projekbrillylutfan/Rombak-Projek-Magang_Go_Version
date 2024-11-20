package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
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

func (service *LokasiServiceImpl) FindAllLokasiService(ctx context.Context) (responses []*web.LokasiCreateOrUpdate) {
	lokasis := service.LokasiRepository.FindAllLokasiRepo(ctx)

	for _, lokasi := range lokasis {
		responses = append(responses, &web.LokasiCreateOrUpdate{
			Nama:   lokasi.Nama,
			Alamat: lokasi.Alamat,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (service *LokasiServiceImpl) FindByIdLokasiService(ctx context.Context, id int64) *web.LokasiModel {
	lokasi, err := service.LokasiRepository.FindByIdLokasiRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return &web.LokasiModel{
		ID:     lokasi.ID,
		Nama:   lokasi.Nama,
		Alamat: lokasi.Alamat,
		CreatedAt: lokasi.CreatedAt.String(),
		UpdateAt:  lokasi.UpdateAt.String(),
	}
}

func (service *LokasiServiceImpl) UpdateLokasiService(ctx context.Context, lokasi *web.LokasiCreateOrUpdate, id int64) *web.LokasiCreateOrUpdate {
	configuration.Validate(lokasi)

	_, err := service.LokasiRepository.FindByIdLokasiRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	lokasis := &domain.Lokasi{
		ID:     id,
		Nama:   lokasi.Nama,
		Alamat: lokasi.Alamat,
	}
	service.LokasiRepository.UpdateLokasiRepo(ctx, lokasis)
	return &web.LokasiCreateOrUpdate{
		Nama:   lokasi.Nama,
		Alamat: lokasi.Alamat,
	}
}

func (service *LokasiServiceImpl) DeleteLokasiService(ctx context.Context, id int64) {
	result, err := service.FindByIdLokasiRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.DeleteLokasiRepo(ctx, result)
}