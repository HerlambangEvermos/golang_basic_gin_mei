package routes

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
	positions := []models.Position{}
	config.DB.Preload("Department").Find(&positions)
	getPositionsResponse := []models.GetPositionResponse{}

	for _, p := range positions {

		pos := models.GetPositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
			Department: models.DepartmentResponse{
				Name: p.Department.Name,
				Code: p.Department.Code,
			},
		}

		getPositionsResponse = append(getPositionsResponse, pos)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    getPositionsResponse,
	})
}

func GetPositionsByID(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.Preload("Department").First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	getPositionResponse := models.GetPositionResponse{
		ID:   position.ID,
		Name: position.Name,
		Code: position.Code,
		Department: models.DepartmentResponse{
			Name: position.Department.Name,
			Code: position.Department.Code,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    getPositionResponse,
	})
}

func PostPositions(c *gin.Context) {
	var position models.Position
	c.BindJSON(&position)

	// insert with GORM
	insert := config.DB.Create(&position)

	if insert.Error != nil {
		log.Printf(insert.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": insert.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    position,
		"message": "success post position",
	})
}

func PutPositions(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	c.BindJSON(&position)

	update := config.DB.Model(&position).Where("id = ?", id).Updates(&position)
	if update.Error != nil {
		log.Printf(update.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": update.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Successfully",
		"data":    position,
	})
}

func DeletePositions(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&position, id)

	c.JSON(200, gin.H{
		"message": "Delete Successfully",
	})

}
