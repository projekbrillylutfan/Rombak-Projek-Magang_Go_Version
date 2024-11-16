package impl_service

import (
	"context"
	"strconv"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
}

type UserServiceImpl struct {
	repository.UserRepository
}

func ConversionError (id string) int64 {
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(exception.ConversionError{
			Message: err.Error(),
		})
	}

	return idInt64
}

func (service *UserServiceImpl) CreateUserService(ctx context.Context, user *web.UserCreate) *web.UserCreate {
	configuration.Validate(user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	users := &domain.User{
		Nama: user.Nama,
		Jabatan: user.Jabatan,
		Username: user.Username,
		Password: string(hashedPassword),
		Role: "USER",
	}
	service.UserRepository.CreateUserRepo(ctx, users)
	user.Password = ""
	return user
}

func (service *UserServiceImpl) FindByIdUserService(ctx context.Context, id int64) *web.UserModel {
	users, err := service.UserRepository.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return &web.UserModel{
		ID: users.ID,
		Nama: users.Nama,
		Jabatan: users.Jabatan,
		Username: users.Username,
	}
}