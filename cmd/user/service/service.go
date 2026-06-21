package service

import (
	"context"
	"userfc/cmd/user/repository"
	"userfc/models"
)

type UserService struct{
	UserRepository *repository.UserRepository
}


func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (svc *UserService) FindUserByEmail (ctx context.Context, email string) (*models.User, error) {
	return svc.UserRepository.FindByEmail(ctx, email)
}

func (svc *UserService) GetUserByID (ctx context.Context, id int64) (*models.User, error) {
	return svc.UserRepository.GetById(ctx, id)
}

func (svc *UserService) CreateUser (ctx context.Context, user *models.User) (int64, error) {
	return svc.UserRepository.InsertNewUser(ctx, user)
}