package loaders

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func dbConfig(ac AppConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		ac.DbHost,
		ac.DbPort,
		ac.DbUser,
		ac.DbName,
		ac.DbPassword,
		ac.DbSslMode,
	)
}

func DBConnect (ac AppConfig) *gorm.DB {
	db, err := gorm.Open("postgres", dbConfig(ac))
	if err != nil {
		fmt.Println("Cannot connect to the database")
		panic(err)
	}

	return db
}