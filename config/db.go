package config

import (
	"log"

	"golang_basic_gin_mei/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_basic_sql_1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("error : %v", err)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Department{}, &models.Position{}, &models.Employee{},
		&models.Inventory{}, &models.Archive{}, &models.EmployeeInventory{}, &models.User{})

	// DB.Create(&models.Department{
	// 	Name: "Human Resources",
	// 	Code: "HRD",
	// 	Positions: []models.Position{
	// 		{Name: "Manager HR", Code: "MHR"},
	// 		{Name: "Staf HR", Code: "SHR"},
	// 	},
	// })

	// DB.Create(&models.Employee{
	// 	Name:       "Herlambang",
	// 	Address:    "Cawang",
	// 	Email:      "herlambang@mail.com",
	// 	PositionID: 1,
	// })

	// DB.Create(&models.Employee{
	// 	Name:       "Heru",
	// 	Address:    "Purbalingga",
	// 	Email:      "heru@mail.com",
	// 	PositionID: 1,
	// })

	// DB.Create(&models.Employee{
	// 	Name:       "Fadlan",
	// 	Address:    "Purbalingga",
	// 	Email:      "fadlan@mail.com",
	// 	PositionID: 2,
	// })

	// DB.Create(&models.Employee{
	// 	Name:       "ruslan",
	// 	Address:    "jaktim",
	// 	Email:      "ruslan@mail.com",
	// 	PositionID: 2,
	// })
}
