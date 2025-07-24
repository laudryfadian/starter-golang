package usecase

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/laudryfadian/starter-golang/internal/domain"
	"github.com/laudryfadian/starter-golang/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *entity.User) error {
	existingUser, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	existingUser, err = u.userRepo.GetByUsername(ctx, user.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if cost == 0 {
		cost = 12
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (entity.AuthResponse, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return entity.AuthResponse{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.AuthResponse{}, errors.New("wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return entity.AuthResponse{}, err
	}

	return entity.AuthResponse{Token: tokenString, User: user}, nil
}

func (u *userUsecase) GetProfile(ctx context.Context, userID string) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, userID)
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID string, user *entity.User) error {
	currentUser, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.Email != currentUser.Email {
		existingUser, err := u.userRepo.GetByEmail(ctx, user.Email)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if existingUser != nil {
			return errors.New("email already exists")
		}
	}

	if user.Username != currentUser.Username {
		existingUser, err := u.userRepo.GetByUsername(ctx, user.Username)
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		if existingUser != nil {
			return errors.New("username already exists")
		}
	}

	return u.userRepo.Update(ctx, userID, user)
}

func (u *userUsecase) DeleteUser(ctx context.Context, userID string) error {
	return u.userRepo.Delete(ctx, userID)
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return u.userRepo.GetAll(ctx)
}
