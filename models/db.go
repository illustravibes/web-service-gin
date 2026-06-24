package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("albums.db"),
		&gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database!")
	}

	err = database.AutoMigrate(&Album{})
	if err != nil {
		panic("Failed to Migrate Database!")
	}

	var count int64
	database.Model(&Album{}).Count(&count)
	if count == 0 {
		dummyAlbums := []Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mullygan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}
		for _, album := range dummyAlbums {
			database.Create(&album)
		}
	}
	DB = database
}
