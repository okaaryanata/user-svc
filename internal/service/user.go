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

func (uc *UserService) CreateUser(ctx context.Context, args *domain.UserRequest) (*domain.User, error) {
	return uc.userRepo.CreateUser(ctx, args)
}

func (uc *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return uc.userRepo.GetUsers(ctx)
}

func (uc *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return uc.userRepo.GetUserByID(ctx, id)
}
