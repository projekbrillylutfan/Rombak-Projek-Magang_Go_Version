package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type BupatiRepository interface {
	CreateBupatiRepo(ctx context.Context, bupati *domain.Bupati) *domain.Bupati
	FindAllBupatiRepo(ctx context.Context) []*domain.Bupati
	FindByIdBupatiRepo(ctx context.Context, id int64) (*domain.Bupati, error)
	UpdateBupatiRepo(ctx context.Context, bupati *domain.Bupati) *domain.Bupati
	DeleteBupatiRepo(ctx context.Context, bupati *domain.Bupati)
}