package controllers

import (
	"net/http"
	"strconv"
	"test-back/config"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

func GetTechniqueByID(c *gin.Context) {
	var techniques models.Techniques
	if err := config.DB.First(&techniques, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "techniques not found"})
		return
	}
	c.JSON(http.StatusOK, techniques)
}

// Получение всех элементов с пагинацией, фильтрацией и сортировкой
func GetTechniques(c *gin.Context) {
	var items []models.Techniques
	query := config.DB

	// Фильтрация по имени
	if id := c.Query("id"); id != "" {
		query = query.Where("id = ?", id)
	}

	// Фильтрация по имени (независимо от регистра)
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Фильтрация по типу техники
	if typeID := c.Query("type_id"); typeID != "" {
		typeIDInt, _ := strconv.Atoi(typeID)
		query = query.Where("type_id = ?", typeIDInt)
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
	query.Model(&models.Techniques{}).Count(&total)

	// Получение данных с лимитом и смещением
	query.Limit(limit).Offset(offset).Find(&items)

	// Ответ с результатами и метаданными
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"limit": limit,
		"items": items,
	})
}

func CreateTechniques(c *gin.Context) {
	var input models.Techniques
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateTechniques(c *gin.Context) {
	var item models.Techniques
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	var input models.Techniques
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func DeleteTechniques(c *gin.Context) {
	var Techniques models.Techniques
	if err := config.DB.First(&Techniques, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	// // Логическое удаление (обновляем поле Deleted на true)
	// item.Deleted = true
	// config.DB.Save(&item)

	config.DB.Delete(&Techniques)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted!"})
}
