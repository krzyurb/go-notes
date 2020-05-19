package loaders

import (
	"os"
)

type AppConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
	DbSslMode  string
}

func BuildConfig () AppConfig {
	return AppConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbSslMode:  os.Getenv("DB_SSL_MODE"),
	}
}