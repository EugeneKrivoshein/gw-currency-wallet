package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	JWTSecret     string
	ServerAddress string
}

func LoadConfig(envPath string) (*Config, error) {
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Не удалось загрузить .env: %v. Используются переменные окружения.", err)
	}

	return &Config{
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}, nil
}
