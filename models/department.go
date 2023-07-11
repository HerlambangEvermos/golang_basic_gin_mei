package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	Positions []Position `json:"positions"`
}

type DepartmentResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetDepartmentResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Code      string             `json:"code"`
	Positions []PositionResponse `json:"positions"`
}
