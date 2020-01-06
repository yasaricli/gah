package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(c *gin.Context) {
	var model RegisterStruct
	c.BindJSON(&model)

	c.JSON(http.StatusOK, gin.H{"user": model.Password})
}
