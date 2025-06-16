package routes

import (
	"net/http"
	"strconv"

	"github.com/abhik555/EventsAPI/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	UserID := context.GetString("userID")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventbyID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = event.Register(UserID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Registered for event."})
}

func cancelRegisterForEvent(context *gin.Context) {
	UserID := context.GetString("userID")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventbyID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = event.CancelRegistration(UserID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Unregistered for event."})

}
