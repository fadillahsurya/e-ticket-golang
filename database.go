package main

import (
	"fmt"
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=123 dbname=eticket port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	fmt.Println("✅ Database connected")

	// Auto migrate
	DB.AutoMigrate(&User{}, &Terminal{})
}
