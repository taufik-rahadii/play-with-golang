package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
)

type Database struct {
}

var DB *gorm.DB

func Connect() {
	var err error

	println(DB_HOST)
	dsn := "host=localhost user=postgres password=password dbname=loyalty port=5432 sslmode=disable Timezone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error")
		log.Fatalln(err.Error())
	}

	log.Print("Database Connection Open")
}
