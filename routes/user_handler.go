package routes

import (
	"net/http"

	"github.com/abhik555/EventsAPI/models"
	"github.com/abhik555/EventsAPI/utils"
	"github.com/gin-gonic/gin"
)

func singup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"stacktrace": err})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"stacktrace": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"stacktrace": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password.", "lit": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"stacktrace": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
