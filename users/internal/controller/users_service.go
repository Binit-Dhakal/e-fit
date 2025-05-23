package controller

import (
	"context"

	"github.com/Binit-Dhakal/e-fit/users/internal/model"
	"github.com/Binit-Dhakal/e-fit/users/internal/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.repo.FindByEmail(ctx, email)
}

func (u *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return u.repo.Create(ctx, user)
}
