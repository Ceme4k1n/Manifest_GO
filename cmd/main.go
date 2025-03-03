package main

import (
	"manifest_go/api/routes"
	"manifest_go/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	port := config.GetEnv("PORT", "8080")

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":" + port)
}
