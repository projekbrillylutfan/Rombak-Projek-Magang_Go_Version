package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type BupatiService interface {
	CreateBupatiService(ctx context.Context, bupati *web.BupatiCreateOrUpdate) *web.BupatiCreateOrUpdate
}