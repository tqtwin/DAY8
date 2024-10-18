package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres" // Sử dụng driver PostgreSQL
	"gorm.io/gorm"
)

// Get the correct database connection string based on the environment
func getDBConfigByEnv(env string) string {
	var user, password, host, port, name string

	switch env {
	case "dev":
		user = os.Getenv("DEV_DB_USER")
		password = os.Getenv("DEV_DB_PASSWORD")
		host = os.Getenv("DEV_DB_HOST")
		port = os.Getenv("DEV_DB_PORT")
		name = os.Getenv("DEV_DB_NAME")
	case "qc":
		// Cấu hình QC nếu cần
	case "prod":
		// Cấu hình PROD nếu cần
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	// Sửa chuỗi kết nối cho PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, name)
	println(dsn)
	return dsn
}

func ConnectDB() (*gorm.DB, error) {
	env := os.Getenv("ENV")
	dsn := getDBConfigByEnv(env)

	// Kết nối với PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// In thông báo khi kết nối thành công
	log.Println("Connected to the database successfully!")

	return db, nil
}
