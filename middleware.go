package gah

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (a *GinAuth) AuthRequiredMiddleware(c *gin.Context) {

	userString := c.GetHeader("X-User-Id")
	authToken := c.GetHeader("X-Auth-Token")
	userID, userIDErr := primitive.ObjectIDFromHex(userString)

	if userIDErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	_, authErr := a.AuthBackEnd.GetUserByToken(userString, authToken)

	if authErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	// set userID
	c.Set("userID", userID)

	// Pass on to the next-in-chain
	c.Next()
}
