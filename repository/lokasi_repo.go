package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type LokasiRepository interface {
	CreateLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) *domain.Lokasi
}