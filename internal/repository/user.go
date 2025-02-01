package repository

import (
	"context"
	"errors"
	"time"

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

func (u *UserRepository) CreateUser(ctx context.Context, args *domain.UserRequest) (*domain.User, error) {
	user := &domain.User{
		Name:      args.Name,
		CreatedAt: time.Now().UnixMicro(),
		UpdatedAt: time.Now().UnixMicro(),
	}

	res := u.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (u *UserRepository) GetUsers(ctx context.Context, args *domain.GetUserRequest) ([]domain.User, error) {
	var users []domain.User
	offset := (args.Page - 1) * args.Size
	u.db.Order("created_at desc").
		Limit(args.Size).Offset(offset).
		Find(&users)
	for i := range users {
		users[i].CreatedAt *= 1000
		users[i].UpdatedAt *= 1000
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return users, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	res := u.db.Where(&domain.User{ID: id}).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (u *UserRepository) GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	var user domain.User
	res := u.db.Where(&domain.User{Name: name}).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
