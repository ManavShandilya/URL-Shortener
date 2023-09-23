package database

import (
	"fmt"
	"os"

	"github.com/Garry028/url-shortener/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

// type connection struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DBName   string
// }

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}
	// connInfo := connection{
	// 	Host:     os.Getenv("POSTGRES_URL"),
	// 	Port:     os.Getenv("POSTGRES_PORT"),
	// 	User:     os.Getenv("POSTGRES_USER"),
	// 	Password: os.Getenv("POSTGRES_PASSWORD"),
	// 	DBName:   os.Getenv("POSTGRES_DB"),
	// }

	DATABASE, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_DB")), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Connected to database\n")
	}
	err = DATABASE.AutoMigrate(&model.Goly{})
	if err != nil {
		fmt.Printf("Error migrating database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Migrated database\n")
	}

}
// func connToString(info connection) string {
// 	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		info.Host, info.Port, info.User, info.Password, info.DBName)
// }
