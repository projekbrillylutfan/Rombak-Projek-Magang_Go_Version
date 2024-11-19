package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type BupatiService interface {
	CreateBupatiService(ctx context.Context, bupati *web.BupatiCreateOrUpdate) *web.BupatiCreateOrUpdate
	FindAllBupatiService(ctx context.Context) []*web.BupatiFindAll
	FindByIdService(ctx context.Context, id int64)*web.BupatiModel
	UpdateBupatiService(ctx context.Context, bupati *web.BupatiCreateOrUpdate, id int64) *web.BupatiCreateOrUpdate
}