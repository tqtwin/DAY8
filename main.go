// main.go
package main

import (
	"day8/config"
	_ "day8/docs" // Đây là cần thiết cho tài liệu Swagger
	"day8/models" // Import model User và Post
	"day8/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Tạo router mới
	router := gin.Default()

	// Load biến môi trường từ file .env
	err := config.LoadEnv()
	if err != nil {
		panic("Failed to load .env file")
	}

	// Kết nối đến cơ sở dữ liệu PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Tự động tạo bảng cho model User và Post
	err = db.AutoMigrate(&models.User{}, &models.Post{}) // Tạo bảng cho model User và Post
	if err != nil {
		panic("Failed to migrate database!")
	}

	// Kết nối đến Redis (nếu có)
	redisCli, err := config.ConnectRedis()
	if err != nil {
		panic("Failed to connect to Redis!")
	}

	// Thiết lập các route
	routes.SetupRoutes(router, db, redisCli)

	// Thiết lập Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Chạy server trên cổng 8080
	router.Run(":8080")
}
