package main

import (
	"test-back/config"
	"test-back/controllers"
	"test-back/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// Подключаем базу данных
	config.ConnectDatabase()
	// Автоматическая миграция
	config.DB.AutoMigrate(&models.Item{})

	r.GET("/items", controllers.GetItems)
	r.POST("/items", controllers.CreateItem)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)

	r.Run(":8080")
}
