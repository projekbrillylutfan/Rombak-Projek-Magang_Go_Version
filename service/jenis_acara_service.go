package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type JenisAcaraService interface {
	CreateJenisAcaraService(ctx context.Context, JenisAcara *web.JenisAcaraCreateOrUpdate) *web.JenisAcaraCreateOrUpdate
}