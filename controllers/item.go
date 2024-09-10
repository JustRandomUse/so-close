package controllers

import (
	"net/http"
	"test-back/config"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

// Получение всех элементов с пагинацией и фильтрацией
func GetItems(c *gin.Context) {
	var items []models.Item
	query := config.DB

	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	//page := c.DefaultQuery("page", "1")
	//limit := c.DefaultQuery("limit", "10")

	var total int64
	query.Model(&models.Item{}).Count(&total)

	//query.Limit(limit).Offset((page - 1) * limit).Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": items,
	})
}

// Добавление нового элемента
func CreateItem(c *gin.Context) {
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// Обновление элемента
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

// Удаление элемента
func DeleteItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found!"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted!"})
}
