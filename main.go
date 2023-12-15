package main

import (
	"example.com/events-rest-api/db"
	"example.com/events-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegesterRoutes(server)

	server.Run(":8080")
}
