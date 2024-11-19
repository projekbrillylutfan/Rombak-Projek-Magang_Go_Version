package impl_repo

import (
	"context"
	"errors"

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

func (repo *BupatiRepositoryImpl)FindAllBupatiRepo(ctx context.Context) []*domain.Bupati {
	var bupatis []*domain.Bupati
	repo.WithContext(ctx).Find(&bupatis)
	return bupatis
}

func (repo *BupatiRepositoryImpl)FindByIdBupatiRepo(ctx context.Context, id int64) (*domain.Bupati, error) {
	var bupati *domain.Bupati
	result:= repo.DB.WithContext(ctx).Unscoped().Where("id_bupati = ?", id).First(&bupati)
	if result.RowsAffected == 0 {
		return nil, errors.New("bupati not found")
	}
	return bupati, nil
}