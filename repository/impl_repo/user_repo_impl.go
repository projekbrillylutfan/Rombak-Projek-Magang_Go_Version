package impl_repo

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"gorm.io/gorm"
)

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func (repo *UserRepositoryImpl) CreateUserRepo(ctx context.Context, user *domain.User) *domain.User {
	err := repo.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}