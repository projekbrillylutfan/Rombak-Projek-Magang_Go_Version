package impl_service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func NewUserServiceImpl(userRepository *repository.UserRepository, config *configuration.Config, redisService *RedisService, mailDialer *gomail.Dialer) service.UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
		Config:         *config,
		RedisService:   redisService,
		MailDialer:     mailDialer,
	}
}

type UserServiceImpl struct {
	repository.UserRepository
	configuration.Config
	*RedisService
	MailDialer *gomail.Dialer
}

// user role 'SUPERADMIN', 'ADMIN', 'USER'

func (service *UserServiceImpl) CreateUserService(ctx context.Context, user *web.UserCreateOrUpdate) *web.UserCreateOrUpdate {
	configuration.Validate(user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	users := &domain.User{
		Nama:     user.Nama,
		Jabatan:  user.Jabatan,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     user.Role,
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
		ID:       users.ID,
		Nama:     users.Nama,
		Jabatan:  users.Jabatan,
		Username: users.Username,
	}
}

func (service *UserServiceImpl) FindAllUserService(ctx context.Context) (responses []*web.UserModel) {
	users := service.UserRepository.FindAllUserRepo(ctx)

	for _, user := range users {
		responses = append(responses, &web.UserModel{
			ID:       user.ID,
			Nama:     user.Nama,
			Jabatan:  user.Jabatan,
			Username: user.Username,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (service *UserServiceImpl) UpdateUserService(ctx context.Context, user *web.UserCreateOrUpdate, id int64) *web.UserCreateOrUpdate {
	configuration.Validate(user)

	_, err := service.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	fmt.Println("isi payload di service update 1 : ", user)
	users := domain.User{
		ID:       id,
		Nama:     user.Nama,
		Jabatan:  user.Jabatan,
		Username: user.Username,
		Password: string(hashedPassword),
	}

	fmt.Println("isi payload di service update 2 : ", users)
	result := service.UpdateUserRepo(ctx, &users)

	return &web.UserCreateOrUpdate{
		Nama:     result.Nama,
		Jabatan:  result.Jabatan,
		Username: result.Username,
		Password: result.Password,
	}
}

func (service *UserServiceImpl) DeleteUserService(ctx context.Context, id int64) {
	result, err := service.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.DeleteUserRepo(ctx, result)
}

func (service *UserServiceImpl) RegisterUserService(ctx context.Context, user *web.UserRegister) *web.UserRegister {
	// validasi user
	configuration.Validate(user)
	// simpan user ke Database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	users := &domain.User{
		Nama:     user.Nama,
		Jabatan:  user.Jabatan,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     "USER",
	}
	service.RegisterUserRepo(ctx, users)

	// Generate token untuk verifikasi email
	tokenBytes := make([]byte, 16)
	if _, err := rand.Read(tokenBytes); err != nil {
		exception.PanicLogging(err)
	}
	token := hex.EncodeToString(tokenBytes)

	// Simpan token di Redis dengan waktu kadaluarsa (misal: 15 menit)
	tokenKey := fmt.Sprintf("email_verification:%s", token)
	err = service.RedisService.Set(ctx, tokenKey, users.ID, time.Minute*15)
	exception.PanicLogging(err)

	// Kirim email verifikasi
	verificationURL := fmt.Sprintf("http://localhost:3000/api/verify-email?token=%s", token)
	message := gomail.NewMessage()
	message.SetHeader("From", "your-email@example.com")
	message.SetHeader("To", users.Email)
	message.SetHeader("Subject", "Email Verification")
	message.SetBody("text/plain", fmt.Sprintf("Hi %s, please verify your email using this link: %s", users.Nama, verificationURL))
	err = service.MailDialer.DialAndSend(message)
	exception.PanicLogging(err)

	user.Password = ""
	return user
}

func (service *UserServiceImpl) VerifyEmailService(ctx context.Context, token string) error {
	// Cek token di Redis
	tokenKey := fmt.Sprintf("email_verification:%s", token)
	userID, err := service.RedisService.Get(ctx, tokenKey)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "invalid or expired token",
		})
	}

	// Update status user menjadi terverifikasi
	id, _ := strconv.ParseInt(userID, 10, 64)
	err = service.UserRepository.MarkUserEmailVerified(ctx, id)
	if err != nil {
		return err
	}

	// Hapus token dari Redis setelah digunakan
	service.RedisService.Del(ctx, tokenKey)
	return nil
}

func (service *UserServiceImpl) Authentication(ctx context.Context, model *web.UserLogin) string {
	configuration.Validate(model)
	user, err := service.UserRepository.AuthenticationRepo(ctx, model.Username)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(model.Password))

	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "usename or password is incorrect",
		})
	}

	tokenJwtResult := utils.GenerateTokenJWT(user.Username, user.Role, service.Config)

	return tokenJwtResult
}

func (service *UserServiceImpl) ForgotPasswordService(ctx context.Context, email string) error {
	user, err := service.FindByEmail(ctx, email)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	// Generate token reset password
	tokenBytes := make([]byte, 16)
	if _, err := rand.Read(tokenBytes); err != nil {
		return err
	}
	token := hex.EncodeToString(tokenBytes)

	// Simpan token di Redis dengan waktu kadaluarsa (contoh: 15 menit)
	tokenKey := fmt.Sprintf("reset_password:%s", token)
	err = service.RedisService.Set(ctx, tokenKey, user.ID, time.Minute*15)
	if err != nil {
		return err
	}

	// Kirim email reset password
	resetURL := fmt.Sprintf("http://localhost:3000/api/reset-password?token=%s", token)
	message := gomail.NewMessage()
	message.SetHeader("From", "your-email@example.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Reset Your Password")
	message.SetBody("text/plain", fmt.Sprintf("Hi %s, reset your password using this link: %s", user.Nama, resetURL))
	err = service.MailDialer.DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) ResetPasswordService(ctx context.Context, token, newPassword string) error {
	// Cek token di Redis
	tokenKey := fmt.Sprintf("reset_password:%s", token)
	userID, err := service.RedisService.Get(ctx, tokenKey)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "invalid or expired token",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	// Update password di database
	id, _ := strconv.ParseInt(userID, 10, 64)
	err = service.UserRepository.UpdatePassword(ctx, id, string(hashedPassword))
	exception.PanicLogging(err)

	// Hapus token dari Redis setelah digunakan
	service.RedisService.Del(ctx, tokenKey)
	return nil
}
