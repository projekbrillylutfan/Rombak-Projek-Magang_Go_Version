package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

func NewJenisAcaraServiceImpl(jenisAcaraRepository *repository.JenisAcaraRepository) service.JenisAcaraService{
	return &JenisAcaraServiceImpl{
		JenisAcaraRepository: *jenisAcaraRepository,
	}
}

type JenisAcaraServiceImpl struct {
	repository.JenisAcaraRepository
}

func(service *JenisAcaraServiceImpl)CreateJenisAcaraService(ctx context.Context, JenisAcara *web.JenisAcaraCreateOrUpdate) *web.JenisAcaraCreateOrUpdate{
	configuration.Validate(JenisAcara)
	JenisAcaras := &domain.JenisAcara{
		NamaJenisAcara: JenisAcara.NamaJenisAcara,
	}
	service.JenisAcaraRepository.CreateJenisAcaraRepo(ctx, JenisAcaras)
	return JenisAcara
}

func (service *JenisAcaraServiceImpl)FindAllJenisAcaraService(ctx context.Context) (responses []*web.JenisAcaraCreateOrUpdate){
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