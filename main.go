package main

import (
	"golang_basic_gin_mei/config"
	"golang_basic_gin_mei/midlewares"
	"golang_basic_gin_mei/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	r.GET("/", getHome)

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := v1.Group("/departments").Use(midlewares.Auth())
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentByID)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}

		position := v1.Group("/positions")
		{
			position.GET("/", routes.GetPositions)
			position.GET("/:id", routes.GetPositionsByID)
			position.POST("/", routes.PostPositions)
			position.PUT("/:id", routes.PutPositions)
			position.DELETE("/:id", routes.DeletePositions)
		}

		employee := v1.Group("/employees").Use(midlewares.Auth())
		{
			employee.GET("/", routes.GetEmployees)
		}

		inventory := v1.Group("/inventories").Use(midlewares.Auth())
		{
			inventory.GET("/", routes.GetInventory)
			inventory.GET("/:id", routes.GetInventoryByID)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}

		rental := v1.Group("/rental").Use(midlewares.Auth())
		{
			rental.GET("/", routes.GetRental)
			rental.GET("/employee/:id", routes.GetRentalByEmployeeID)
			rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
			rental.POST("/employee", routes.PostRentalByEmployee)
		}

	}

	r.Run() // listen and serve on localhost:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
