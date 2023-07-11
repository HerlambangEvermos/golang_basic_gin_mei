package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	// Employees   []*Employee `gorm:"many2many:employees_inventories"`
	Employees []EmployeeInventory `json:"employees"`
}

type RequestInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	Archive              ResponseArchive
}

type EmployeeInventory struct {
	gorm.Model
	EmployeeID  uint `json:"employee_id"`
	Employee    Employee
	InventoryID uint `json:"inventory_id"`
	Inventory   Inventory
	Description string `json:"description"`
}
