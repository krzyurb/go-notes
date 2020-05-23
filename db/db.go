package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"notesapp/loaders"
)

func dbConfig(ac loaders.AppConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		ac.DbHost,
		ac.DbPort,
		ac.DbUser,
		ac.DbName,
		ac.DbPassword,
		ac.DbSslMode)
}

func BuildConnection (ac loaders.AppConfig) *gorm.DB {
	db, err := gorm.Open("postgres", dbConfig(ac))
	if err != nil {
		fmt.Println("Cannot connect to the database")
		panic(err)
	}

	return db
}