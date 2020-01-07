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

// Login User handler
func LoginHandler(c *gin.Context) {
	var validate *validator.Validate
	var body LoginStruct

	// parse body post data
	c.BindJSON(&body)

	// Validate LoginStruct form!
	validate = validator.New()
	validateError := validate.Struct(body)

	if validateError != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessageResponse(validateError.Error()))
		return
	}

	user, err := utils.GetUserEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessageResponse("Unauthorized"))
		return
	}

	log.Println(user, err)

	c.JSON(http.StatusOK, utils.SuccessDataResponse(gin.H{
		"user": user,
	}))
}
