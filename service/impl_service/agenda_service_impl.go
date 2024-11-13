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

func NewAgendaServiceImpl(agendaRepository *repository.AgendaRepository, bupatiRepository *repository.BupatiRepository, jenisAcaraRepository *repository.JenisAcaraRepository, lokasiRepository *repository.LokasiRepository) service.AgendaService {
	return &AgendaServiceImpl{
		AgendaRepository: *agendaRepository,
		BupatiRepository: *bupatiRepository,
		JenisAcaraRepository: *jenisAcaraRepository,
		LokasiRepository: *lokasiRepository,
	}
}

type AgendaServiceImpl struct {
	repository.AgendaRepository
	repository.BupatiRepository
	repository.JenisAcaraRepository
	repository.LokasiRepository
}

func (service *AgendaServiceImpl) CreateAgendaService(ctx context.Context, agenda *web.AgendaCreateOrUpdate) *web.AgendaCreateOrUpdate {
	configuration.Validate(agenda)
	// cek bupati 
	_, err := service.BupatiRepository.FindByIdBupatiRepo(ctx, agenda.IDBupati)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	// cek jenis acara
	_, err = service.FindByIdJenisAcaraRepo(ctx, agenda.IDJenisAcara)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	// cek lokasi
	_, err = service.FindByIdLokasiRepo(ctx, agenda.IDLokasi)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	agendas := &domain.Agenda{
		IDBupati: agenda.IDBupati,
		IDJenisAcara: agenda.IDJenisAcara,
		IDLokasi: agenda.IDLokasi,
		NamaAgenda: agenda.NamaAgenda,
		Deskripsi: agenda.Deskripsi,
		TanggalMulai: agenda.TanggalMulai,
		TanggalSelesai: agenda.TanggalSelesai,
	}
	service.AgendaRepository.CreateAgendaRepo(ctx, agendas)
	return agenda
}

func (service *AgendaServiceImpl) FindAllAgendaService(ctx context.Context) (responses []*web.AgendaCreateOrUpdate) {
	agedas := service.AgendaRepository.FindAllAgendaRepo(ctx)

	for _, agenda := range agedas {
		responses = append(responses, &web.AgendaCreateOrUpdate{
			IDBupati: agenda.IDBupati,
			IDJenisAcara: agenda.IDJenisAcara,
			IDLokasi: agenda.IDLokasi,
			NamaAgenda: agenda.NamaAgenda,
			Deskripsi: agenda.Deskripsi,
			TanggalMulai: agenda.TanggalMulai,
			TanggalSelesai: agenda.TanggalSelesai,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (service *AgendaServiceImpl)FindByIdAgendaService(ctx context.Context, id int64) *web.AgendaModel {
	agenda, err := service.AgendaRepository.FindByIdAgendaRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return &web.AgendaModel{
		ID: agenda.IDAgenda,
		NamaAgenda: agenda.NamaAgenda,
		Bupati: web.ConvertBupatiToModel(&agenda.Bupati),
		Deskripsi: agenda.Deskripsi,
		Lokasi: web.ConvertLokasiToModel(&agenda.Lokasi),
		JenisAcara: web.ConvertJenisAcaraToModel(&agenda.JenisAcara),
		TanggalMulai: agenda.TanggalMulai,
		TanggalSelesai: agenda.TanggalSelesai,
	}
}

// update repository di agenda
func (service *AgendaServiceImpl)UpdateAgendaService(ctx context.Context, agenda *web.AgendaCreateOrUpdate, id int64) *web.AgendaCreateOrUpdate {
	configuration.Validate(agenda)

	// cek id agenda
	_, err := service.AgendaRepository.CheckIDAgendaRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	
	// cek bupati 
	_, err = service.BupatiRepository.FindByIdBupatiRepo(ctx, agenda.IDBupati)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	// cek jenis acara
	_, err = service.FindByIdJenisAcaraRepo(ctx, agenda.IDJenisAcara)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	// cek lokasi
	_, err = service.FindByIdLokasiRepo(ctx, agenda.IDLokasi)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	agendas := &domain.Agenda{
		IDAgenda: id,
		IDBupati: agenda.IDBupati,
		IDJenisAcara: agenda.IDJenisAcara,
		IDLokasi: agenda.IDLokasi,
		NamaAgenda: agenda.NamaAgenda,
		Deskripsi: agenda.Deskripsi,
		TanggalMulai: agenda.TanggalMulai,
		TanggalSelesai: agenda.TanggalSelesai,
	}

	result := service.AgendaRepository.UpdateAgendaRepo(ctx, agendas)
	return &web.AgendaCreateOrUpdate{
		NamaAgenda: result.NamaAgenda,
		IDBupati: result.IDBupati,
		Deskripsi: result.Deskripsi,
		IDLokasi: result.IDLokasi,
		IDJenisAcara: result.IDJenisAcara,
		TanggalMulai: result.TanggalMulai,
		TanggalSelesai: result.TanggalSelesai,
	}
}

func (service *AgendaServiceImpl)DeleteAgendaService(ctx context.Context, id int64) {
	result, err := service.AgendaRepository.CheckIDAgendaRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.AgendaRepository.DeleteAgendaRepo(ctx, result)
}