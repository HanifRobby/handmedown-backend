package controllers

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(context *gin.Context) {

	db := config.DB

	var users []models.User

	// Query untuk mencari data produk
	err := db.Find(&users)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating HTTP Response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    users,
	})
}

type ProfileResponse struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	NoTelp string `json:"noTelp"`
	Alamat string `json:"alamat"`
}

func GetProfile(context *gin.Context) {
	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID ke uint
	userIDUint, ok := userID.(uint)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID type"})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", userIDUint).First(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting users profile"})
		return
	}

	profileResponse := ProfileResponse{
		Nama:   user.Nama,
		Email:  user.Email,
		NoTelp: user.NoTelp,
		Alamat: user.Alamat,
	}

	context.JSON(http.StatusOK, gin.H{
		"data": profileResponse,
	})
}
