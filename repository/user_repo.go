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
	UpdateUserByEmailRepo(email string, updates map[string]any) error
}