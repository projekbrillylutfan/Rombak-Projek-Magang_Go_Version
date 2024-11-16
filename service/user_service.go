package service

import (
	"context"

	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

type UserService interface {
	CreateUserService(ctx context.Context, user *web.UserCreateOrUpdate) *web.UserCreateOrUpdate
	FindByIdUserService(ctx context.Context, id int64) *web.UserModel
	FindAllUserService(ctx context.Context) []*web.UserModel
	UpdateUserService(ctx context.Context, user *web.UserCreateOrUpdate, id int64) *web.UserCreateOrUpdate
}