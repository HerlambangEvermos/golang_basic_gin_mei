package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	InventoryID uint
}

type ResponseArchive struct {
	ArchiveName        string `json:"archive_name"`
	ArchiveDescription string `json:"archive_description"`
}
