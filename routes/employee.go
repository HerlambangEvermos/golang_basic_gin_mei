package routes

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	employees := []models.Employee{}
	config.DB.Preload("Position").Find(&employees)

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    employees,
	})
}
