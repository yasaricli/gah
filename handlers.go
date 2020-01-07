package gah

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// LoginHandler Login User handler
func LoginHandler(c *gin.Context) {
	var validate *validator.Validate
	var body LoginStruct

	// parse body post data
	c.BindJSON(&body)

	// Validate LoginStruct form!
	validate = validator.New()
	validateError := validate.Struct(body)

	if validateError != nil {
		c.JSON(http.StatusBadRequest, ErrorMessageResponse(validateError.Error()))
		return
	}

	user, err := GetUserEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorMessageResponse("Unauthorized"))
		return
	}

	log.Println(user, err)

	c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
		"user": user,
	}))
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
		c.JSON(http.StatusBadRequest, ErrorMessageResponse(validateError.Error()))
		return
	}

	// check if the user has already registered.
	_, userError := GetUserEmail(body.Email)

	if userError != nil {
		insertedUser := InsertUser(body.Email, body.Password)

		c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
			"_id":   insertedUser.ID,
			"email": insertedUser.Email,
		}))

		return
	}

	c.JSON(http.StatusConflict, ErrorMessageResponse("Username already exists."))
}
