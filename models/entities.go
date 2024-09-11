package models

// import (
// 	"gorm.io/gorm"
// )

type EmployeeEquipments struct {
	// gorm.Model
	ID          int `json:"id"`
	EmployeesID int `json:"employees_id"`
	TechniqueID int `json:"technique_id"`
}

type Techniques struct {
	// gorm.Model
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TypeID int    `json:"type_id"`
	// Archived bool   `json:"archived"`
	// Deleted  bool   `json:"deleted"`
}

type TechniqueTypes struct {
	// gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Employees struct {
	// gorm.Model
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Office int    `json:"office"`
	// Archived bool   `json:"archived"`
	// Deleted  bool   `json:"deleted"`
}
