package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/abhik555/EventsAPI/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err1 := models.GetAllEvents()
	if err1 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := models.GetEventbyID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, event)

}

func creatEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"stacktrace": err})
		return
	}

	event.UserID = context.GetString("userID")
	event.CreatedAt = time.Now()
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event": event})
}

func updateEvents(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := models.GetEventbyID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	UserID := context.GetString("userID")

	if event.UserID != UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "This operation is not permitted."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updatedEvent.ID = int(id)
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	event, err := models.GetEventbyID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	UserID := context.GetString("userID")

	if event.UserID != UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "This operation is not permitted."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfuly"})
}
