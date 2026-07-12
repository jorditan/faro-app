package internal

import (
	"net/http"

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

	users := router.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.FindAll)
		users.GET("/:id", userHandler.FindByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}
}
