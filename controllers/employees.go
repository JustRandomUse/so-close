package controllers

import (
	"net/http"
	"strconv"
	"test-back/config"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

func GetEmployeeByID(c *gin.Context) {
	var employee models.Employees
	if err := config.DB.First(&employee, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// Получение всех элементов с пагинацией и фильтрацией
func GetEmployees(c *gin.Context) {
	var Employees []models.Employees
	query := config.DB

	// Фильтрация по имени
	if id := c.Query("id"); id != "" {
		query = query.Where("id = ?", id)
	}

	// Фильтрация по имени
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Фильтрация по офису
	if office := c.Query("office"); office != "" {
		query = query.Where("office = ?", office)
	}

	// Фильтрация по статусу архивности
	if archived := c.Query("archived"); archived != "" {
		archivedBool, _ := strconv.ParseBool(archived)
		query = query.Where("archived = ?", archivedBool)
	}

	// Фильтрация по статусу удаления
	if deleted := c.Query("deleted"); deleted != "" {
		deletedBool := deleted == "true"
		query = query.Where("deleted = ?", deletedBool)
	}

	// Пагинация
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var total int64
	query.Model(&models.Employees{}).Count(&total)

	// Получение данных с лимитом и смещением
	query.Limit(limit).Offset(offset).Find(&Employees)

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
		"items": Employees,
	})
}

// Добавление нового элемента
func CreateEmployees(c *gin.Context) {
	var input models.Employees
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание нового элемента
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// Обновление элемента
func UpdateEmployees(c *gin.Context) {
	var Employees models.Employees
	if err := config.DB.First(&Employees, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	var input models.Employees
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление существующего элемента
	config.DB.Model(&Employees).Updates(input)
	c.JSON(http.StatusOK, Employees)
}

// Удаление элемента (soft delete)
func DeleteEmployees(c *gin.Context) {
	var Employees models.Employees
	if err := config.DB.First(&Employees, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	// // Логическое удаление (обновляем поле Deleted на true)
	// Employees.Deleted = true
	// config.DB.Save(&Employees)

	config.DB.Delete(&Employees)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted!"})
}
