package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pandakn/gofiber-crud/config"
	"github.com/pandakn/gofiber-crud/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {

	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Todo{})

	DB = DbInstance{
		Db: db,
	}
}
