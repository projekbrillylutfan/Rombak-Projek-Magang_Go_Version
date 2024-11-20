package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type JenisAcaraRepository interface {
	CreateJenisAcaraRepo(ctx context.Context, JenisAcara *domain.JenisAcara) *domain.JenisAcara
}