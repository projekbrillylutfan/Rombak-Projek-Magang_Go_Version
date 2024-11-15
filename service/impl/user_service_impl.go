package impl

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
}

type UserServiceImpl struct {
	repository.UserRepository
}

func (service *UserServiceImpl) CreateUserService(ctx context.Context, user *web.UserCreate) *web.UserCreate {
	configuration.Validate(user)
	users := &domain.User{
		Nama: user.Nama,
		Jabatan: user.Jabatan,
		Username: user.Username,
		Password: user.Password,
		Role: "USER",
	}
	service.UserRepository.CreateUserRepo(ctx, users)
	return user
}