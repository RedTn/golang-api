package main

import (
	"fmt"
	"golang-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
