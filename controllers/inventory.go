package controllers

import (
	"net/http"
	"test-back/config"

	"github.com/gin-gonic/gin"
)

// Модель для представления данных инвентаря
type InventoryView struct {
	EmployeeName       string `json:"employee_name"`
	EmployeeID         int    `json:"employee_id"`
	TechniqueType      string `json:"technique_type"`
	TechniqueName      string `json:"technique_name"`
	TechniqueTabNumber int    `json:"technique_tab_number"`
}

// Получение данных инвентаря по номеру офиса
func GetInventoryByOffice(c *gin.Context) {
	var inventory []InventoryView
	office := c.Param("office")

	// Проверяем, что номер офиса передан
	if office == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Office number is required"})
		return
	}

	// Запрос к представлению inventory_view с фильтром по офису
	if err := config.DB.Raw(`
        SELECT * FROM inventory_view
        WHERE employee_id IN (SELECT id FROM employees WHERE office = ?)
    `, office).Scan(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"office":    office,
		"inventory": inventory,
	})
}
