package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type UserRepository interface {
	CreateUserRepo(ctx context.Context, user *domain.User) *domain.User
	FindByIdUserRepo(ctx context.Context, id int64) (*domain.User, error)
	FindAllUserRepo(ctx context.Context) []*domain.User
}