package routes

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetInventory(c *gin.Context) {
	inventory := []models.Inventory{}
	config.DB.Preload(clause.Associations).Find(&inventory)

	resInventories := []models.ResponseInventory{}

	for _, inv := range inventory {
		resInv := models.ResponseInventory{
			InventoryName:        inv.Name,
			InventoryDescription: inv.Description,
			Archive: models.ResponseArchive{
				ArchiveName:        inv.Archive.Name,
				ArchiveDescription: inv.Archive.Description,
			},
		}

		resInventories = append(resInventories, resInv)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get inventory",
		"data":    resInventories,
	})
}

func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")

	var inventory models.Inventory
	data := config.DB.Preload(clause.Associations).First(&inventory, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	resInv := models.ResponseInventory{
		InventoryName:        inventory.Name,
		InventoryDescription: inventory.Description,
		Archive: models.ResponseArchive{
			ArchiveName:        inventory.Archive.Name,
			ArchiveDescription: inventory.Archive.Description,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resInv,
	})
}

func PostInventory(c *gin.Context) {
	var reqInventory models.RequestInventory
	c.BindJSON(&reqInventory)

	// insert with GORM
	inventory := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.ArchiveDescription,
		}}

	config.DB.Create(&inventory)

	c.JSON(http.StatusCreated, gin.H{
		"data":    reqInventory,
		"message": "success post inventory",
	})

}

func PutInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory
	data := config.DB.First(&inventory, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	var reqInventory models.RequestInventory
	c.BindJSON(&reqInventory)
	inventoryUpdate := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.ArchiveDescription,
		},
	}

	config.DB.Model(&inventory).Where("id = ?", id).Updates(&inventoryUpdate)

	c.JSON(200, gin.H{
		"message": "Update Successfully",
		"data":    reqInventory,
	})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory

	data := config.DB.First(&inventory, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&inventory, id)

	c.JSON(200, gin.H{
		"message": "Delete Successfully",
	})

}
