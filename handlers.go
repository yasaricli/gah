package gah

import (
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

	user, err := GetUserByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorMessageResponse("Unauthorized"))
		return
	}

	pwd := []byte(body.Password)
	pwdMatch := ComparePasswords(user.Password, pwd)

	if pwdMatch {

		// Add a new auth token
		token := InsertHashedLoginToken(user.ID)

		// XXX: STATUS OK
		c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
			"authToken": token,
			"userId":    user.ID,
		}))

		return
	}

	// user not check pass
	c.JSON(http.StatusUnauthorized, ErrorMessageResponse("Unauthorized"))
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
	_, userError := GetUserByEmail(body.Email)

	if userError != nil {
		insertedUser := CreateUser(body.Email, body.Password)

		c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
			"_id":   insertedUser.ID,
			"email": insertedUser.Email,
		}))

		return
	}

	c.JSON(http.StatusConflict, ErrorMessageResponse("Username already exists."))
}
