package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zykunov/bankAPI/handlers"

	_ "github.com/zykunov/bankAPI/docs"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/")
	{
		apiGroup.POST("/accounts", handlers.AddAccount)            // Добавление аккаунта
		apiGroup.POST("/accounts/:id/deposit", handlers.Deposit)   // Пополнение баланса
		apiGroup.POST("/accounts/:id/withdraw", handlers.Withdraw) // Вывод денех
		apiGroup.GET("/accounts/:id/balance", handlers.GetBalance) // Получение баланса по id пользователя
	}

	return router
}
