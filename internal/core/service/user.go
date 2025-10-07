package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func InitUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	// hash password
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return &domain.User{}, err
	}

	user.Password = hashedPwd

	// create user
	res, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}

	// return response
	return res, nil
}
