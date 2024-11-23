package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type UserRepository interface {
	CreateUserRepo(ctx context.Context, user *domain.User) *domain.User
	FindByIdUserRepo(ctx context.Context, id int64) (*domain.User, error)
	FindAllUserRepo(ctx context.Context) []*domain.User
	UpdateUserRepo(ctx context.Context, user *domain.User) *domain.User
	DeleteUserRepo(ctx context.Context, user *domain.User)
	RegisterUserRepo(ctx context.Context, user *domain.User) *domain.User
	AuthenticationRepo(ctx context.Context, username string) (*domain.User, error)
	MarkUserEmailVerified(ctx context.Context, userID int64) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdatePassword(ctx context.Context, userID int64, hashedPassword string) error
}