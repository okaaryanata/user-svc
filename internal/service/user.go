package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/okaaryanata/user/internal/domain"
	"github.com/okaaryanata/user/internal/repository"
)

type (
	UserService struct {
		userRepo *repository.UserRepository
	}
)

func NewUserService(
	userRepo *repository.UserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (uc *UserService) CreateUser(ctx context.Context, args *domain.UserRequest) (*domain.User, error) {
	user, err := uc.userRepo.GetUserByName(ctx, args.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("user with name: %s already exist", args.Name)
	}

	return uc.userRepo.CreateUser(ctx, args)
}

func (uc *UserService) GetUsers(ctx context.Context, args *domain.GetUserRequest) ([]domain.User, error) {
	return uc.userRepo.GetUsers(ctx, args)
}

func (uc *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return uc.userRepo.GetUserByID(ctx, id)
}
