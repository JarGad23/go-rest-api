package main

import (
	"github.com/JarGad23/go-rest-api/db"
	"github.com/JarGad23/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
