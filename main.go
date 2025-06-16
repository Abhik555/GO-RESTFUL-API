package main

import (
	"github.com/abhik555/EventsAPI/db"
	"github.com/abhik555/EventsAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080 <-- IP address
}
