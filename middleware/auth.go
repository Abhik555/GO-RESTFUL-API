package middleware

import (
	"net/http"

	"github.com/Abhik555/GO-RESTFUL-API/utils"
	"github.com/gin-gonic/gin"
)

func Auththenticate(context *gin.Context) {
	authToken := context.Request.Header.Get("Authorization")

	if authToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"stacktrace": "Invalid login token"})
		return
	}

	userID, err := utils.VerifyToken(authToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"stacktrace": "Invalid login token"})
		return
	}

	context.Set("userID", userID)
	context.Next()

}
