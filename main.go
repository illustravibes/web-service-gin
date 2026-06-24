package main

import (
	"example/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbum)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	router.Run("localhost:8080")
}
