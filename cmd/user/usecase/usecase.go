package usecase

import (
	"context"
	"errors"
	"time"
	"userfc/cmd/user/service"
	"userfc/infrastructure/log"
	"userfc/models"
	"userfc/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserUsecase struct {
	UserService *service.UserService
	Secret      string
}

func NewUserUsecase(userService *service.UserService, secret string) *UserUsecase {
	return &UserUsecase{
		UserService: userService ,
		Secret: secret,
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

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := uc.UserService.FindUserByEmail(ctx, email)
	if err != nil {
		log.Logger.Errorf("token.SignedString got error: %v", err)
		return "", err
	}

	if user == nil {
		return "", errors.New("Wrong Email")
	}

	isMatch := utils.Verify(user.Password, password)

	if !isMatch {
		return "", errors.New("Wrong Password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(uc.Secret))
	if err != nil {
		log.Logger.Errorf("token.SignedString got error: %v", err)
		return "", err
	}

	return tokenString, nil
}