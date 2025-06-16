package routes

import (
	"github.com/abhik555/EventsAPI/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authGroup := server.Group("/")
	authGroup.Use(middleware.Auththenticate)
	authGroup.POST("/events", creatEvent)
	authGroup.PUT("/events/:id" , updateEvents)
	authGroup.DELETE("/events/:id" , deleteEvent)
	authGroup.POST("/events/:id/register" , registerForEvent)
	authGroup.DELETE("/events/:id/register" , cancelRegisterForEvent)

	server.POST("/signup" , singup)
	server.POST("/login", login)
}
