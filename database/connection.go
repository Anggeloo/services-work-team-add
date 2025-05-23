package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST=aws-postgres-microservices.cxfclnhjpr44.us-east-1.rds.amazonaws.com
	DB_PORT=5432
	DB_USER=postgres
	DB_PASSWORD=AWSPServices0099
	DB_NAME=db_work_team
	

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	DB = database
	fmt.Println("Database connected successfully")
}
