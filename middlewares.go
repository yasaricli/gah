package gah

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthRequiredMiddleware // AuthRequiredMiddleware() middleware just in the "authorized"
func AuthRequiredMiddleware(c *gin.Context) {
	userID, userIDErr := primitive.ObjectIDFromHex(c.GetHeader("X-User-Id"))
	authToken := c.GetHeader("X-Auth-Token")

	if userIDErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	_, authErr := GetUserByToken(userID, authToken)

	if authErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessageResponse("You must be logged in to do this."))
		return
	}

	// set userID
	c.Set("userID", userID)

	// Pass on to the next-in-chain
	c.Next()
}
