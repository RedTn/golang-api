package main

import (
	"fmt"
	"golang-api/src/handlers"

	"golang-api/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)

	// runs this code async/starts a new thread
	go func() {
		err := router.Run("localhost:8080")
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	utils.PrintMessage("Util called")

	// halts the program so it does not terminate
	select {}
}
