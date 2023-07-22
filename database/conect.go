package database

import (
	"log"
	"os"

	"github.com/KayoRonald/go-fiber-jwt-test/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance


func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}
	log.Print("Connected Successfully to Database")
	
	// db.Logger = logger.Default.LogMode(logger.Info)
	log.Print("Running Migrations")
	db.AutoMigrate(&models.User{})

	Database = DbInstance{
		Db: db,
	}
}