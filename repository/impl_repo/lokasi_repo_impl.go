package impl_repo

import (
	"context"
	"errors"

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

func (repo *LokasiRepositoryImpl) FindByIdLokasiRepo(ctx context.Context, id int64) (*domain.Lokasi, error) {
	var lokasi *domain.Lokasi
	result := repo.DB.WithContext(ctx).Unscoped().Where("id_lokasi = ?", id).First(&lokasi)
	if result.RowsAffected == 0 {
		return &domain.Lokasi{}, errors.New("lokasi not found")
	}
	return lokasi, nil
}

func (repo *LokasiRepositoryImpl) UpdateLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) *domain.Lokasi {
	err := repo.DB.WithContext(ctx).Where("id_lokasi = ?", lokasi.ID).Updates(&lokasi).Error
	exception.PanicLogging(err)
	return lokasi
}

func (repo *LokasiRepositoryImpl) DeleteLokasiRepo(ctx context.Context, lokasi *domain.Lokasi) {
	err := repo.DB.WithContext(ctx).Delete(&lokasi).Error
	exception.PanicLogging(err)
}