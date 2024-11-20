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

func NewJenisAcaraServiceImpl(jenisAcaraRepository *repository.JenisAcaraRepository) service.JenisAcaraService {
	return &JenisAcaraServiceImpl{
		JenisAcaraRepository: *jenisAcaraRepository,
	}
}

type JenisAcaraServiceImpl struct {
	repository.JenisAcaraRepository
}

func (service *JenisAcaraServiceImpl) CreateJenisAcaraService(ctx context.Context, JenisAcara *web.JenisAcaraCreateOrUpdate) *web.JenisAcaraCreateOrUpdate {
	configuration.Validate(JenisAcara)
	JenisAcaras := &domain.JenisAcara{
		NamaJenisAcara: JenisAcara.NamaJenisAcara,
	}
	service.JenisAcaraRepository.CreateJenisAcaraRepo(ctx, JenisAcaras)
	return JenisAcara
}

func (service *JenisAcaraServiceImpl) FindAllJenisAcaraService(ctx context.Context) (responses []*web.JenisAcaraCreateOrUpdate) {
	jenisAcaras := service.FindAllJenisAcaraRepo(ctx)
	for _, jenisAcara := range jenisAcaras {
		responses = append(responses, &web.JenisAcaraCreateOrUpdate{
			NamaJenisAcara: jenisAcara.NamaJenisAcara,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (service *JenisAcaraServiceImpl) FindByIdJenisAcaraService(ctx context.Context, id int64) *web.JenisAcaraModel {
	jenisAcara, err := service.FindByIdJenisAcaraRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return &web.JenisAcaraModel{
		ID:             jenisAcara.ID,
		NamaJenisAcara: jenisAcara.NamaJenisAcara,
		CreatedAt:      jenisAcara.CreatedAt.String(),
		UpdateAt:       jenisAcara.UpdateAt.String(),
	}
}

func(service *JenisAcaraServiceImpl) UpdateJenisAcaraService(ctx context.Context, JenisAcara *web.JenisAcaraCreateOrUpdate, id int64) *web.JenisAcaraCreateOrUpdate {
	configuration.Validate(JenisAcara)

	_, err := service.FindByIdJenisAcaraRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	jenisAcara := &domain.JenisAcara{
		ID:             id,
		NamaJenisAcara: JenisAcara.NamaJenisAcara,
	}
	service.UpdateJenisAcaraRepo(ctx, jenisAcara)
	return JenisAcara
}

func (service *JenisAcaraServiceImpl) DeleteJenisAcaraService(ctx context.Context, id int64) {
	result, err := service.FindByIdJenisAcaraRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.DeleteJenisAcaraRepo(ctx, result)
}
