package handlers

import (
	"net/http"

	"../utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v10"
)

type RegisterStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

// RegisterHandler Gin register handler
func RegisterHandler(c *gin.Context) {
	var validate *validator.Validate
	var body RegisterStruct

	// parse body post data
	c.BindJSON(&body)

	// Validate LoginStruct form!
	validate = validator.New()
	validateError := validate.Struct(body)

	if validateError != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessageResponse(validateError.Error()))
		return
	}

	// check if the user has already registered.
	_, userError := utils.GetUserEmail(body.Email)

	if userError != nil {
		insertedUser := utils.InsertUser(body.Email, body.Password)

		c.JSON(http.StatusOK, utils.SuccessDataResponse(gin.H{
			"_id":   insertedUser.ID,
			"email": insertedUser.Email,
		}))

		return
	}

	c.JSON(http.StatusConflict, utils.ErrorMessageResponse("Username already exists."))
}
