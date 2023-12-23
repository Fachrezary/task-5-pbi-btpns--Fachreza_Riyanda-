// controllers/photo_controller.go

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"projectapi/app"
	"projectapi/database"
)

// CreatePhoto handles the creation of a new photo.
func CreatePhoto(c *gin.Context) {
	var photo app.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating photo"})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

// GetPhotos retrieves all photos.
func GetPhotos(c *gin.Context) {
	var photos []app.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto updates photo information.
func UpdatePhoto(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	var photo app.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating photo"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

// DeletePhoto deletes a photo by ID.
func DeletePhoto(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	if err := database.DB.Delete(&app.Photo{}, photoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}
