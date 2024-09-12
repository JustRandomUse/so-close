package controllers

import (
	"net/http"
	"strconv"
	"test-back/config"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

// func GetEmployeeEquipmentByID(c *gin.Context) {
// 	var employee_equipment models.EmployeeEquipments
// 	if err := config.DB.First(&employee_equipment, c.Param("id")).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "employee_equipment not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, employee_equipment)

// }

type EmployeeEquipmentWithOffice struct {
	EmployeeName  string `json:"employee_name"`
	EmployeeID    int    `json:"employee_id"`
	Office        int    `json:"office"`
	TechniqueID   int    `json:"technique_id"`
	TechniqueName string `json:"technique_name"`
}

func GetEmployeeEquipmentByID(c *gin.Context) {
	var employeeEquipment models.EmployeeEquipments
	if err := config.DB.First(&employeeEquipment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee_equipment not found"})
		return
	}

	// Запрос данных сотрудника по employees_id
	var employee models.Employees
	if err := config.DB.First(&employee, employeeEquipment.EmployeesID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Запрос данных техники по technique_id
	var technique models.Techniques
	if err := config.DB.First(&technique, employeeEquipment.TechniqueID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Technique not found"})
		return
	}

	// Формируем ответ с данными сотрудника, офиса и техники
	response := EmployeeEquipmentWithOffice{
		EmployeeName:  employee.Name,
		EmployeeID:    employee.ID,
		Office:        employee.Office,
		TechniqueID:   technique.ID,
		TechniqueName: technique.Name,
	}

	c.JSON(http.StatusOK, response)
}

// Получение всех элементов с пагинацией и фильтрацией
func GetEmployeeEquipments(c *gin.Context) {
	var items []models.EmployeeEquipments
	query := config.DB

	// Фильтрация по имени
	if id := c.Query("id"); id != "" {
		query = query.Where("id = ?", id)
	}

	// Фильтрация по EmployeeID
	if employeeID := c.Query("employees_id"); employeeID != "" {
		employeeIDInt, _ := strconv.Atoi(employeeID)
		query = query.Where("employees_id = ?", employeeIDInt)
	}

	// Фильтрация по TechniqueID
	if techniqueID := c.Query("technique_id"); techniqueID != "" {
		techniqueIDInt, _ := strconv.Atoi(techniqueID)
		query = query.Where("technique_id = ?", techniqueIDInt)
	}

	// Пагинация
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var total int64
	query.Model(&models.EmployeeEquipments{}).Count(&total)

	// Получение данных с лимитом и смещением
	query.Limit(limit).Offset(offset).Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
		"items": items,
	})
}

func CreateEmployeeEquipments(c *gin.Context) {
	var input models.EmployeeEquipments
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateEmployeeEquipments(c *gin.Context) {
	var item models.EmployeeEquipments
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	var input models.EmployeeEquipments
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func DeleteEmployeeEquipments(c *gin.Context) {
	var item models.EmployeeEquipments
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted!"})
}
