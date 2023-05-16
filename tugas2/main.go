package main

import (
	"log"
	"tugas2/handler"
	"tugas2/orders"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&orders.Order{}, &orders.Item{})

	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderHandler := handler.NewOrdersHandler(orderService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.GET("/orders", orderHandler.FindAll)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.POST("/orders", orderHandler.Save)
	api.PUT("/orders/:id", orderHandler.Update)
	api.DELETE("/orders/:id", orderHandler.Delete)

	router.Run()
}
