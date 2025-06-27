package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://neondb_owner:npg_n3UTZ4isPhJu@ep-plain-grass-a8cpxetm-pooler.eastus2.azure.neon.tech/go-server?sslmode=require"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("Connected to database")
}
