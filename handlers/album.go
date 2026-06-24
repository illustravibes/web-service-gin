package handlers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbum(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	if err := models.DB.First(&album, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&newAlbum).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := models.DB.First(&album, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
		return
	}

	var input models.Album
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := models.DB.First(&album, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
		return
	}

	models.DB.Delete(&album)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album has been deleted successfully."})
}
