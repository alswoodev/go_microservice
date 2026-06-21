package usecase

import (
	"context"
	"userfc/cmd/user/service"
	"userfc/infrastructure/log"
	"userfc/models"
	"userfc/utils"
)

type UserUsecase struct {
	UserService *service.UserService
}

func NewUserUsecase(userService *service.UserService) *UserUsecase {
	return &UserUsecase{
		UserService: userService ,
	}
}

func (uc *UserUsecase) FindUserByEmail (ctx context.Context , email string) (*models.User, error) {
	return uc.UserService.FindUserByEmail(ctx , email)
}

func (uc *UserUsecase) GetUserByID (ctx context.Context , id int64) (*models.User, error) {
	return uc.UserService.GetUserByID(ctx, id)
}

func (uc *UserUsecase) RegisterUser (ctx context.Context, user *models.User) error {
	hashed, err := utils.Hash(user.Password)
	if err != nil {
		log.Logger.Fatalf("Failed to hash: %v", err)
		return err
	}
	
	user.Password = hashed
	_, err = uc.UserService.CreateUser(ctx, user)
	if err != nil {
		log.Logger.Fatalf("Failed to Insert User: %v", err)
		return err
	}

	return nil
}