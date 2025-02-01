package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/okaaryanata/user/internal/domain"
)

type (
	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	u.db.Order("created_at desc").Find(&users)
	for i := range users {
		users[i].CreatedAt *= 1000
		users[i].UpdatedAt *= 1000
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return users, nil
}
