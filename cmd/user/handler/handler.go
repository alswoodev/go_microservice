package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{

}

// It returns pointer of UserHandler when successful.
// Otherwise, nil pointer of UserHandler will be returned.
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Ping (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status":"ok"})
}