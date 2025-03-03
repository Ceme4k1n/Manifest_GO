package routes

import (
	"manifest_go/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.Ping)
		api.GET("/getUsers", handlers.GetUsers)
	}
}
