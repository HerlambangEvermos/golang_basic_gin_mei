package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
}

type PositionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetPositionResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Department DepartmentResponse
}
