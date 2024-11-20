package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type LokasiRepository interface {
	CreateLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) *domain.Lokasi
	FindAllLokasiRepo(ctx context.Context) []*domain.Lokasi
	FindByIdLokasiRepo(ctx context.Context, id int64) (*domain.Lokasi, error)
	UpdateLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) *domain.Lokasi
	DeleteLokasiRepo(ctx context.Context, lokasi *domain.Lokasi)
}