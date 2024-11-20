package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type LokasiService interface {
	CreateLokasiService(ctx context.Context, lokasi *web.LokasiCreateOrUpdate) *web.LokasiCreateOrUpdate
	FindAllLokasiService(ctx context.Context) []*web.LokasiCreateOrUpdate
}