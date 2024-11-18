package impl_repo

import (
	"context"
	"errors"
	"fmt"

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

func (repo *UserRepositoryImpl) FindByIdUserRepo(ctx context.Context, id int64) (*domain.User, error) {
	var user *domain.User
	result := repo.DB.WithContext(ctx).Unscoped().Where("id_user = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return &domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (repo *UserRepositoryImpl)FindAllUserRepo(ctx context.Context) []*domain.User {
	var users []*domain.User
	repo.WithContext(ctx).Find(&users)
	return users
}

func (repo *UserRepositoryImpl) UpdateUserRepo(ctx context.Context, user *domain.User) *domain.User {
	err := repo.DB.WithContext(ctx).Where("id_user = ?", user.ID).Updates(&user).Error
	exception.PanicLogging(err)
	fmt.Println("isi payload di repo update ", user)
	return user
}

func (repo *UserRepositoryImpl) DeleteUserRepo(ctx context.Context, user *domain.User) {
	err := repo.DB.WithContext(ctx).Delete(&user).Error
	exception.PanicLogging(err)
}

func (repo *UserRepositoryImpl) RegisterUserRepo(ctx context.Context, user *domain.User) *domain.User {
	err := repo.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (repo *UserRepositoryImpl) AuthenticationRepo(ctx context.Context, username string) (*domain.User, error) {
	var user *domain.User
	result := repo.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		return &domain.User{}, errors.New("user not found")
	}
	return user, nil
}