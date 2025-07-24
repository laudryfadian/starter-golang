package domain

import (
	"context"

	"github.com/laudryfadian/starter-golang/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, id string, user *entity.User) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
}
