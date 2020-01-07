package gah

import (
	"github.com/gin-gonic/gin"
)

// ErrorMessageResponse Request response object ready for errors.
func ErrorMessageResponse(message string) gin.H {
	return gin.H{
		"status": "error",
		"data": gin.H{
			"message": message,
		},
	}
}

// SuccessMessageResponse Request response object ready for success.
func SuccessMessageResponse(message string) gin.H {
	return gin.H{
		"status": "success",
		"data": gin.H{
			"message": message,
		},
	}
}

// SuccessDataResponse gives a successful request response.
func SuccessDataResponse(data gin.H) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
	}
}
