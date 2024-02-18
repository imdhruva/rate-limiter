package main

import (
	"fmt"
	"net/http"
	"os"
	"rate-limiter/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := InitRoutes()
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(0)
	}
}

func InitRoutes() *gin.Engine {

	router := gin.Default()
	router.GET("/hello", middleware.RateLimit, sayOK)
	router.GET("/health", middleware.RateLimit, sayHealthy)
	return router
}

func sayOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func sayHealthy(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
