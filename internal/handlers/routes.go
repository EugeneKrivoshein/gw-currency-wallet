package handlers

import (
	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/storages"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, storage storages.Storage) {
	authHandler := NewAuthHandler(storage)

	auth := router.Group("/auth")
	{
		auth.POST("/api/v1/register", authHandler.Register)
		auth.POST("/api/v1/login", authHandler.Login)
		auth.POST("/api/v1/logout", authHandler.Logout) // Добавлен маршрут для логаута
	}
}
