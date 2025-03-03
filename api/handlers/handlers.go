package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PONG"})
}

func GetUsers(c *gin.Context) {
	users := []string{"Alice", "Bob", "Charlie"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
