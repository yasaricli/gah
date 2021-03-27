package gah

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

func (a *GinAuth) LoginHandler(c *gin.Context) {

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

	user, err := a.AuthBackEnd.GetUserByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorMessageResponse("Unauthorized"))
		return
	}

	pwd := []byte(body.Password)
	pwdMatch := a.AuthBackEnd.ComparePasswords(user.Password, pwd)

	if pwdMatch {
		token := a.AuthBackEnd.InsertHashedLoginToken(user.ID.String())
		c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
			"authToken": token,
			"userId":    user.ID,
		}))

		return
	}

	c.JSON(http.StatusUnauthorized, ErrorMessageResponse("Unauthorized"))
}

func (a *GinAuth) RegisterHandler(c *gin.Context) {
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
	_, userError := a.AuthBackEnd.GetUserByEmail(body.Email)

	if userError != nil {
		insertedUser := a.AuthBackEnd.CreateUser(body.Email, body.Password)

		c.JSON(http.StatusOK, SuccessDataResponse(gin.H{
			"_id":   insertedUser.ID,
			"email": insertedUser.Email,
		}))

		return
	}

	c.JSON(http.StatusConflict, ErrorMessageResponse("Username already exists."))
}
