package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type UserService interface {
	CreateUserService(ctx context.Context, user *web.UserCreate) *web.UserCreate
}