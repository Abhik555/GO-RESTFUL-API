package main

import (
	"github.com/Abhik555/GO-RESTFUL-API/db"
	"github.com/Abhik555/GO-RESTFUL-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.SetTrustedProxies(nil)
	gin.SetMode(gin.ReleaseMode)

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080 <-- IP address
}
