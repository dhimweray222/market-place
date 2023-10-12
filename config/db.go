package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}
var DB *gorm.DB
func Connect(){
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	// Create the database connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	DB=db
	for _, model := range RegisterModels(){
		err = DB.Debug().AutoMigrate(model.Model)
	}
}
