package domain

import (
	"context"

	"github.com/laudryfadian/starter-golang/internal/entity"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, email, password string) (entity.AuthResponse, error)
	GetProfile(ctx context.Context, userID string) (*entity.User, error)
	UpdateProfile(ctx context.Context, userID string, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
}
