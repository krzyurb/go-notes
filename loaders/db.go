package loaders

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func DBConnect () *gorm.DB {
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("Cannot connect to the database")
		panic(err)
	}

	return db
}