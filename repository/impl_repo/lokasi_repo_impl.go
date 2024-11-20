package impl_repo

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"gorm.io/gorm"
)

func NewLokasiRepository(DB *gorm.DB) repository.LokasiRepository {
	return &LokasiRepositoryImpl{
		DB: DB,
	}
}

type LokasiRepositoryImpl struct {
	*gorm.DB
}

func (repo *LokasiRepositoryImpl) CreateLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) *domain.Lokasi {
	err := repo.DB.WithContext(ctx).Create(&lokasi).Error
	exception.PanicLogging(err)
	return lokasi
}

func (repo *LokasiRepositoryImpl) FindAllLokasiRepo(ctx context.Context) []*domain.Lokasi {
	var lokasis []*domain.Lokasi
	err := repo.DB.WithContext(ctx).Find(&lokasis).Error
	exception.PanicLogging(err)
	return lokasis
}