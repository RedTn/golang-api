package main

import (
	"golang-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)

	router.Run("localhost:8080")
}
