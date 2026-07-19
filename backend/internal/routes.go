package internal

import (
	"net/http"

	"faro/backend/internal/category"
	"faro/backend/internal/transaction"
	"faro/backend/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	userRepository := user.NewMemoryRepository()
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	categoryRepository := category.NewMemoryRepository()
	categoryService := category.NewService(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)
	transactionRepository := transaction.NewMemoryRepository()
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := transaction.NewHandler(transactionService)

	users := router.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.FindAll)
		users.GET("/:id", userHandler.FindByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	categories := router.Group("/categories")
	{
		categories.POST("", categoryHandler.Create)
		categories.GET("", categoryHandler.FindAll)
		categories.GET("/:id", categoryHandler.FindByID)
		categories.PUT("/:id", categoryHandler.Update)
		categories.DELETE("/:id", categoryHandler.Delete)
	}

	transactions := router.Group("/transactions")
	{
		transactions.POST("", transactionHandler.Create)
		transactions.GET("", transactionHandler.FindAll)
		transactions.GET("/:id", transactionHandler.FindByID)
		transactions.PUT("/:id", transactionHandler.Update)
		transactions.DELETE("/:id", transactionHandler.Delete)
	}
}
