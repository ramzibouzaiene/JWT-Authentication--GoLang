package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error

	dsn := os.Getenv("DBINFO")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	println("database connected")
	if err != nil {
		panic("Database connection failed")
	}
}
