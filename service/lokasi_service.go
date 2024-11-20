package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type LokasiService interface {
	CreateLokasiService(ctx context.Context, lokasi *web.LokasiCreateOrUpdate) *web.LokasiCreateOrUpdate
	FindAllLokasiService(ctx context.Context) []*web.LokasiCreateOrUpdate
	FindByIdLokasiService(ctx context.Context, id int64) *web.LokasiModel
	UpdateLokasiService(ctx context.Context, lokasi *web.LokasiCreateOrUpdate, id int64) *web.LokasiCreateOrUpdate
	DeleteLokasiService(ctx context.Context, id int64)
}