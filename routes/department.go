package routes

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDepartment(c *gin.Context) {
	departments := []models.Department{}
	// config.DB.Find(&departments)
	config.DB.Preload("Positions").Find(&departments)
	// config.DB.Preload(clause.Associations).Find(&departments)

	GetDepartmentResponses := []models.GetDepartmentResponse{}

	for _, d := range departments {
		positions := []models.PositionResponse{}
		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		GetDepartmentResponses = append(GetDepartmentResponses, dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    GetDepartmentResponses,
	})
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")

	var department models.Department
	data := config.DB.Preload("Positions").First(&department, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	positions := []models.PositionResponse{}
	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		}

		positions = append(positions, pos)
	}

	GetDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": GetDepartmentResponse,
	})
}

func PostDepartment(c *gin.Context) {

	// cara post ver x-www-form
	// department := models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// }

	//cara post dengan json
	var department models.Department
	c.BindJSON(&department)

	// insert with GORM
	config.DB.Create(&department)

	c.JSON(http.StatusCreated, gin.H{
		"data":    department,
		"message": "success post department",
	})

}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	// cara ver x-www-form
	// data := config.DB.First(&department, "id = ?", id)
	// if data.Error != nil {
	// 	log.Printf(data.Error.Error())
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"status":  http.StatusNotFound,
	// 		"message": "Data not found",
	// 	})

	// 	return
	// }

	// config.DB.Model(&department).Updates(models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// })

	//cara ver json

	data := config.DB.First(&department, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	c.BindJSON(&department)

	config.DB.Model(&department).Where("id = ?", id).Updates(&department)

	c.JSON(200, gin.H{
		"message": "Update Successfully",
		"data":    department,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	data := config.DB.First(&department, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&department, id)

	c.JSON(200, gin.H{
		"message": "Delete Successfully",
	})

}
