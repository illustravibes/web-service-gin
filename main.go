package main

import (
	"example/web-service-gin/handlers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()

	router.GET("/albums", handlers.GetAlbum)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	router.Run("localhost:8080")
}
