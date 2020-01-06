package handlers

import (
	"log"
	"net/http"

	"../utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v10"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

var validate *validator.Validate

// Login User handler
func LoginHandler(c *gin.Context) {
	var body LoginStruct

	// parse body post data
	c.BindJSON(&body)

	// Validate LoginStruct form!
	validate = validator.New()
	validateError := validate.Struct(body)

	if validateError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data": gin.H{
				"message": validateError.Error(),
			},
		})

		return
	}

	user, err := utils.GetUser(body.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Unauthorized",
			},
		})

		return
	}

	log.Println(user, err)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
