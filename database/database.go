package database

import (
	"drivehino-trivia-restapi/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct{
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb()  {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
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

	// Next we want to use AutoMigrate to create the tables that we need. We pass all our GORM models to AutoMigrate. In this tutorial, we only have one GORM model, which is the Facts model.
	log.Println("running migrations")
    db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}