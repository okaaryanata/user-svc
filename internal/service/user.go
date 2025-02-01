package service

import (
	"context"

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

func (uc *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	users, err := uc.userRepo.GetUsers(ctx)
	if err != nil {
		return users, err
	}

	return users, nil
}
