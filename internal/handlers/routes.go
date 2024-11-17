package handlers

import (
	"github.com/EugeneKrivoshein/gw-currency-wallet/internal/storages"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, storage storages.Storage) {
	authHandler := NewAuthHandler(storage)

	router.POST("/api/v1/register", authHandler.Register)
	router.POST("/api/v1/login", authHandler.Login)
	router.POST("/api/v1/logout", authHandler.Logout)

}
