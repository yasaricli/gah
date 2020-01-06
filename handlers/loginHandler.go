package handlers

import (
	"log"
	"net/http"

	"../utils"
	"github.com/gin-gonic/gin"
)

type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login User handler
func LoginHandler(c *gin.Context) {
	var body LoginStruct

	// parse body post data
	c.BindJSON(&body)

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
