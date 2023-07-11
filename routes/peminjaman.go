package routes

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type ResponseRental struct {
	EmployeeID    uint      `json:"employee_id"`
	InventoryID   uint      `json:"inventory_id"`
	EmployeeName  string    `json:"employee_name"`
	InventoryName string    `json:"inventory_name"`
	Description   string    `json:"description"`
	RentalDate    time.Time `json:"rental_date"`
}

func GetRental(c *gin.Context) {
	EmployeeInventories := []models.EmployeeInventory{}

	config.DB.Preload(clause.Associations).Find(&EmployeeInventories)

	resRentals := []ResponseRental{}

	for _, eInv := range EmployeeInventories {
		resRent := ResponseRental{
			EmployeeID:    eInv.EmployeeID,
			InventoryID:   eInv.InventoryID,
			EmployeeName:  eInv.Employee.Name,
			InventoryName: eInv.Inventory.Name,
			Description:   eInv.Description,
			RentalDate:    eInv.CreatedAt,
		}

		resRentals = append(resRentals, resRent)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental",
		"data":    resRentals,
	})
}

func GetRentalByEmployeeID(c *gin.Context) {
	id := c.Param("id")

	var employeeInventories []models.EmployeeInventory
	data := config.DB.Preload(clause.Associations).Find(&employeeInventories, "employee_id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	resRentals := []ResponseRental{}

	for _, eInv := range employeeInventories {
		resRent := ResponseRental{
			EmployeeID:    eInv.EmployeeID,
			InventoryID:   eInv.InventoryID,
			EmployeeName:  eInv.Employee.Name,
			InventoryName: eInv.Inventory.Name,
			Description:   eInv.Description,
			RentalDate:    eInv.CreatedAt,
		}

		resRentals = append(resRentals, resRent)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resRentals,
	})
}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")

	var employeeInventories []models.EmployeeInventory
	data := config.DB.Preload(clause.Associations).Find(&employeeInventories, "inventory_id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	resRentals := []ResponseRental{}

	for _, eInv := range employeeInventories {
		resRent := ResponseRental{
			EmployeeID:    eInv.EmployeeID,
			InventoryID:   eInv.InventoryID,
			EmployeeName:  eInv.Employee.Name,
			InventoryName: eInv.Inventory.Name,
			Description:   eInv.Description,
			RentalDate:    eInv.CreatedAt,
		}

		resRentals = append(resRentals, resRent)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resRentals,
	})
}

func PostRentalByEmployee(c *gin.Context) {
	reqRental := models.EmployeeInventory{}
	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}
	insert := config.DB.Create(&reqRental)

	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    reqRental,
		"message": "Insert Success",
	})
}
