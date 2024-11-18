package impl_repo

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"gorm.io/gorm"
)

func NewBupatiRepository(DB *gorm.DB) repository.BupatiRepository {
	return &BupatiRepositoryImpl{DB: DB}
}

type BupatiRepositoryImpl struct {
	*gorm.DB
}

func (repo *BupatiRepositoryImpl)CreateBupatiRepo(ctx context.Context, bupati *domain.Bupati) *domain.Bupati {
	err := repo.DB.WithContext(ctx).Create(&bupati).Error
	exception.PanicLogging(err)
	return bupati
}