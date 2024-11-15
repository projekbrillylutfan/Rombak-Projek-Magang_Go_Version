package repository

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/domain"
)

type UserRepository interface {
	CreateUserRepo(ctx context.Context, user *domain.User) *domain.User
}