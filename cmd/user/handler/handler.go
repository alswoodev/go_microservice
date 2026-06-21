package handler

import (
	"net/http"
	"userfc/cmd/user/usecase"
	"userfc/infrastructure/log"
	"userfc/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	UserUsecase *usecase.UserUsecase
}

// It returns pointer of UserHandler when successful.
// Otherwise, nil pointer of UserHandler will be returned.
func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Register (c *gin.Context) {
	var param models.RegisterParameter
	if err := c.ShouldBindJSON(&param); err !=nil {
		log.Logger.Info("Invalid parameter")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if len(param.Password) < 8 {
		log.Logger.Info("Password must longer than 8 characters")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password must longer than 8 characters",
		})
		return
	}

	if param.Password != param.ConfirmPassword {
		log.Logger.Info("Password and Confirm Password not match")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password and Confirm Password not match",
		})
		return
	}

	user, err := h.UserUsecase.FindUserByEmail(c.Request.Context(), param.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Email already exists!",
		})
		return
	}

	err = h.UserUsecase.RegisterUser(c.Request.Context(), &models.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password, // pw will be hashed while running the func.
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully registered!",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var param models.LoginParameter
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Input Parameter",
		})

		return
	}

	token, err := h.UserUsecase.Login(c.Request.Context(), param.Email, param.Password)
	if err != nil {
		log.Logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong Email or Password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *UserHandler) GetMyInfo (c *gin.Context) {
	// context of `user_id` is stored as int
	uid, isExist := c.Get("user_id")
	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user, err := h.UserUsecase.GetUserByID(c.Request.Context(), int64(uid.(int))) // uid -> type assertion -> type conversion
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"email": user.Email,
	})
}


func (h *UserHandler) Ping (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status":"ok"})
}