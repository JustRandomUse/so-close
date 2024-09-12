package controllers

import (
	"net/http"
	"strconv"
	"test-back/config"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

func GetTechniqueTypeByID(c *gin.Context) {
	var techniquetype models.TechniqueTypes
	if err := config.DB.First(&techniquetype, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "techniquetype not found"})
		return
	}
	c.JSON(http.StatusOK, techniquetype)
}

// Получение всех элементов с пагинацией и фильтрацией
func GetTechniqueType(c *gin.Context) {
	var items []models.TechniqueTypes
	query := config.DB

	// Фильтрация по имени
	if id := c.Query("id"); id != "" {
		query = query.Where("id = ?", id)
	}

	// Фильтр по имени (если предоставлено в запросе)
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	// Пагинация
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var total int64
	query.Model(&models.TechniqueTypes{}).Count(&total)

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

func CreateTechniqueType(c *gin.Context) {
	var input models.TechniqueTypes
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateTechniqueType(c *gin.Context) {
	var item models.TechniqueTypes
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	var input models.TechniqueTypes
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func DeleteTechniqueType(c *gin.Context) {
	var item models.TechniqueTypes
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted!"})
}
