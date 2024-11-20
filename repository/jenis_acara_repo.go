package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type JenisAcaraRepository interface {
	CreateJenisAcaraRepo(ctx context.Context, JenisAcara *domain.JenisAcara) *domain.JenisAcara
	FindAllJenisAcaraRepo(ctx context.Context) []*domain.JenisAcara
	FindByIdJenisAcaraRepo(ctx context.Context, id int64) (*domain.JenisAcara, error)
	UpdateJenisAcaraRepo(ctx context.Context, JenisAcara *domain.JenisAcara) *domain.JenisAcara
	DeleteJenisAcaraRepo(ctx context.Context, JenisAcara *domain.JenisAcara)
}