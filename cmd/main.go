package main

import (
	"log"

	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/config"
	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/handlers"
	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/storages"
	postgres "github.com/EugeneKrivoshein/gw-currency-wallet/internal/storages/postgres"
	"github.com/EugeneKrivoshein/gw-currency-wallet/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("config.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	pkg.SetSecret(cfg.JWTSecret)

	dbProvider, err := postgres.NewPostgresProvider(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer func() {
		if err := dbProvider.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с БД: %v", err)
		}
	}()

	if err := postgres.RunMigrations(dbProvider.DB()); err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	var storage storages.Storage = postgres.NewPostgresStorage(dbProvider)

	router := gin.Default()

	handlers.RegisterRoutes(router, storage)

	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
