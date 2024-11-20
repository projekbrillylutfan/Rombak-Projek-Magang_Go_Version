package impl_repo

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"gorm.io/gorm"
)

func NewJenisAcaraRepository(DB *gorm.DB) repository.JenisAcaraRepository {
	return &JenisAcaraRepositoryImpl{
		DB: DB,
	}
}

type JenisAcaraRepositoryImpl struct {
	*gorm.DB
}

func (repo *JenisAcaraRepositoryImpl)CreateJenisAcaraRepo(ctx context.Context, JenisAcara *domain.JenisAcara) *domain.JenisAcara {
	err := repo.DB.WithContext(ctx).Create(&JenisAcara).Error
	exception.PanicLogging(err)
	return JenisAcara
}

func (repo *JenisAcaraRepositoryImpl)FindAllJenisAcaraRepo(ctx context.Context) []*domain.JenisAcara {
	var JenisAcaras []*domain.JenisAcara
	repo.WithContext(ctx).Find(&JenisAcaras)
	return JenisAcaras
}